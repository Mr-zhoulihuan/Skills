package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Tool struct {
	Type     string `json:"type`
	Function Function `json:"function`
}

type Function struct {
	Name        string          `json:"name`
	Description string          `json:"description`
	Parameters  json.RawMessage `json:"parameters`
}

type Message struct {
	Role       string     `json:"role`
	Content    *string    `json:"content`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty`
	ToolCallID string     `json:"tool_call_id,omitempty`
}

type ToolCall struct {
	ID       string `json:"id`
	Type     string `json:"type`
	Function struct {
		Name      string `json:"name`
		Arguments string `json:"arguments`
	} `json:"function`
}

type ChatRequest struct {
	Model      string    `json:"model`
	Messages   []Message `json:"messages`
	Tools      []Tool    `json:"tools,omitempty`
	ToolChoice string    `json:"tool_choice,omitempty`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message`
	} `json:"choices`
	Error *struct {
		Message string `json:"message`
	} `json:"error,omitempty`
}

type ToolHandler func(args json.RawMessage) (string, error)

type WeatherArgs struct {
	City string `json:"city`
}

type CalcArgs struct {
	A  float64 `json:"a`
	B  float64 `json:"b`
	Op string  `json:"op`
}

type SearchArgs struct {
	Query string `json:"query`
}

func handleGetWeather(args json.RawMessage) (string, error) {
	var p WeatherArgs
	if e := json.Unmarshal(args, &p); e != nil {
		return "", fmt.Errorf("bad args: %v", e)
	}
	r := fmt.Sprintf(`{"city":"%s","temp":5,"cond":"sunny"}`, p.City)
	return r, nil
}

func handleCalculate(args json.RawMessage) (string, error) {
	var p CalcArgs
	if e := json.Unmarshal(args, &p); e != nil {
		return "", fmt.Errorf("bad args: %v", e)
	}
	var r float64
	switch p.Op {
	case "+": r = p.A + p.B
	case "-": r = p.A - p.B
	case "*": r = p.A * p.B
	case "/":
		if p.B == 0 { return `{"error":"div0"}`, nil }
		r = p.A / p.B
	default: return "", fmt.Errorf("bad op: %s", p.Op)
	}
	return fmt.Sprintf(`{"result":%v}`, r), nil
}

func handleSearch(args json.RawMessage) (string, error) {
	var p SearchArgs
	if e := json.Unmarshal(args, &p); e != nil {
		return "", fmt.Errorf("bad args: %v", e)
	}
	return fmt.Sprintf(`{"query":"%s","results":["info for %s"]}`, p.Query, p.Query), nil
}

func initTools() ([]Tool, map[string]ToolHandler) {
	tools := []Tool{
		{Type:"function", Function:Function{
			Name:"get_weather",
			Description:"get weather by city name",
			Parameters:json.RawMessage(`{"type":"object","properties":{"city":{"type":"string"}},"required":["city"]}`),
		}},
		{Type:"function", Function:Function{
			Name:"calculate",
			Description:"do math: + - * /",
			Parameters:json.RawMessage(`{"type":"object","properties":{"a":{"type":"number"},"b":{"type":"number"},"op":{"type":"string","enum":["+","-","*","/"]}},"required":["a","b","op"]}`),
		}},
		{Type:"function", Function:Function{
			Name:"search",
			Description:"search web for info",
			Parameters:json.RawMessage(`{"type":"object","properties":{"query":{"type":"string"}},"required":["query"]}`),
		}},
	}
	handlers := map[string]ToolHandler{
		"get_weather": handleGetWeather,
		"calculate":   handleCalculate,
		"search":      handleSearch,
	}
	return tools, handlers
}

func ptr(s string) *string { return &s }

func runFC(question string, tools []Tool, handlers map[string]ToolHandler, choice string) string {
	msgs := []Message{
		{Role: "system", Content: ptr("You are a helpful assistant. Use tools when needed. Answer in Chinese.")},
		{Role: "user", Content: ptr(question)},
	}
	for turn := 0; turn < 15; turn++ {
		body, _ := json.Marshal(ChatRequest{
			Model: os.Getenv("OPENAI_MODEL"),
			Messages: msgs, Tools: tools, ToolChoice: choice,
		})
		url := strings.TrimRight(os.Getenv("OPENAI_BASE_URL"), "/")
		if url == "" { url = "https://api.openai.com/v1" }
		key := os.Getenv("OPENAI_API_KEY")
		if strings.Contains(bodyStr, `"model":""`) {
			return "ERROR: missing OPENAI_MODEL"
		}
		if key == "" { return "ERROR: missing OPENAI_API_KEY" }

		_, _ = url, key
		body, _ = json.Marshal(ChatRequest{
			Model:     func() string { m := os.Getenv("OPENAI_MODEL"); if m == "" { return "gpt-4o" }; return m }(),
			Messages:  msgs,
			Tools:     tools,
			ToolChoice: choice,
		})
		req, _ := http.NewRequest("POST", url+"/chat/completions", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+key)
		resp, e := http.DefaultClient.Do(req)
		if e != nil { return "ERROR: " + e.Error() }
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode != 200 { return "HTTP "+resp.Status+": "+string(b) }

		var cr ChatResponse
		json.Unmarshal(b, &cr)
		if cr.Error != nil { return "API: "+cr.Error.Message }
		msg := cr.Choices[0].Message
		if len(msg.ToolCalls) == 0 {
			if msg.Content != nil { return *msg.Content }
			return ""
		}
		msgs = append(msgs, Message{Role: "assistant", Content: msg.Content, ToolCalls: msg.ToolCalls})

		type res struct { id, val string }
		ch := make(chan res, len(msg.ToolCalls))
		var wg sync.WaitGroup
		for _, tc := range msg.ToolCalls {
			wg.Add(1)
			go func(tc ToolCall) {
				defer wg.Done()
				if h := handlers[tc.Function.Name]; h != nil {
					r, e := h(json.RawMessage(tc.Function.Arguments))
					if e != nil { ch <- res{tc.ID, `{"error":"`+e.Error()+`"}`}; return }
					ch <- res{tc.ID, r}
				} else {
					ch <- res{tc.ID, `{"error":"unknown tool"}`}
				}
			}(tc)
		}
		wg.Wait(); close(ch)
		rm := map[string]string{}
		for r := range ch { rm[r.id] = r.val }
		for _, tc := range msg.ToolCalls {
			msgs = append(msgs, Message{Role: "tool", ToolCallID: tc.ID, Content: ptr(rm[tc.ID])})
		}
	}
	return "ERROR: max turns"
}

func main() {
	tools, handlers := initTools()
	tests := []struct{t, q, c string}{
		{"Single tool: weather", "What is the weather in Beijing?", "auto"},
		{"Multi tool: weather + calc", "Beijing is 5C and Shanghai is 10C, what is the difference?", "auto"},
		{"Force tool_choice=required", "Hello!", "required"},
	}
	for _, tt := range tests {
		fmt.Printf("\n=== %s ===\nQ: %s\n", tt.t, tt.q)
		fmt.Println("A:", runFC(tt.q, tools, handlers, tt.c))
	}
}
