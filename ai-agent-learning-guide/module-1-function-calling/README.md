 # 模块一：工具调用（Function Calling / Tool Use）
 
 ## 是什么
 
 **工具调用**是 AI Agent 的基础能力。LLM 本身无法直接执行代码、查询数据库、发请求、操作文件。通过 Function Calling，LLM 输出一个结构化的函数调用指令（JSON），由你的代码解析并实际执行，再将结果返回给 LLM 继续推理。
 
 ```
 用户提问 ──→ LLM ──→ 输出 {"function": "get_weather", "args": {"city": "北京"}}
                         ↓
                   你的代码执行 get_weather("北京")
                         ↓
                   返回 "北京 25°C 晴" ──→ LLM ──→ 组织最终回答 → 用户
 ```
 
 这不是 LLM 真的"调用"了函数，而是 LLM 以 JSON 格式描述了"应该调用哪个函数、传什么参数"，执行层是你的代码。
 
 ---
 
 ## 核心概念
 
 - **Function Schema**：用 JSON Schema 描述函数的名称、参数、返回值，传给 LLM 的 API
 - **Tool Choice**：控制 LLM 是否必须调用工具（`auto` / `required` / `none`）
 - **Multi-Turn Tool Use**：LLM 连续调用多个工具，逐步完成任务
 - **Parallel Tool Calls**：LLM 一次输出多个函数调用，并发执行
 
 ---
 
 ## 学习路径
 
 ### 第 1 步：跑通第一个 Function Calling（半天）
 
 1. 注册 OpenAI / 国产大模型 API，获取 Key
 2. 写一个 Python 脚本，定义 2 个模拟函数（如 `get_weather`、`calculate`）
 3. 用官方的 Chat Completions API（`tools` 参数）调用 LLM
 4. 解析 `tool_calls` 响应，执行函数，返回结果
 5. 让 LLM 基于函数结果生成最终回答
 
 ```python
 # 最小代码骨架
 response = client.chat.completions.create(
     model="gpt-4o",
     messages=[{"role": "user", "content": "北京和上海哪个冷？"}],
     tools=[{
         "type": "function",
         "function": {
             "name": "get_weather",
             "parameters": {
                 "type": "object",
                 "properties": {
                     "city": {"type": "string"}
                 },
                 "required": ["city"]
             }
         }
     }]
 )
 ```
 
 ### 第 2 步：处理多轮工具调用（半天）
 
 - 模拟一个场景让 LLM 连续调用 2-3 个工具
 - 处理 `Parallel Tool Calls`（一次返回多个 tool_calls）
 - 加入错误处理：工具执行失败时如何反馈给 LLM
 
 ### 第 3 步：对接真实 API（一天）
 
 - 把模拟函数换成真实的外部 API（天气 API、搜索 API、数据库查询）
 - 处理认证、限流、超时
 - 增加工具调用的安全校验（哪些工具可以调用、参数白名单）
 
 ### 第 4 步：用框架包装（半天）
 
 - 用 LangChain 的 `@tool` 装饰器或 `tool()` 函数重新实现
 - 对比手写和框架的差异
 
 ---
 
 ## 资源推荐
 
 | 资源 | 类型 | 链接 |
 |---|---|---|
 | OpenAI Function Calling 官方指南 | 文档 | https://platform.openai.com/docs/guides/function-calling |
 | OpenAI Cookbook - Function Calling 示例 | 代码 | https://cookbook.openai.com/examples/how_to_call_functions_with_chat_models |
 | Azure OpenAI Function Calling | 文档 | https://learn.microsoft.com/azure/ai-services/openai/how-to/function-calling |
 | LangChain Tools 概念 | 文档 | https://python.langchain.com/docs/concepts/tools/ |
 | Anthropic Tool Use 文档 | 文档 | https://docs.anthropic.com/en/docs/build-with-claude/tool-use |
 | 通义千问 Function Calling 文档 | 文档 | https://help.aliyun.com/zh/model-studio/function-calling |
 
 ---
 
 ## 检验标准
 
 - [ ] 能用原生 API 跑通一个完整的工具调用流程
 - [ ] 能处理工具调用失败的情况（重试 / 降级 / 报错）
 - [ ] 能实现并发的平行工具调用
 - [ ] 能口头说清楚 Function Calling 的原理（什么在客户端做什么在服务端做）
 - [ ] 能在 LangChain 中用 `@tool` 定义一个工具并调用
 
 ---
 
 ## 常见误区
 
 - ❌ **认为 LLM "会"调用函数** — 实际上 LLM 只输出 JSON 描述，执行是你的事
 - ❌ **一次性定义太多工具** — 工具越多，LLM 选错工具的概率越大。从 2-3 个开始
 - ❌ **忽略参数校验** — LLM 可能生成非法参数，需要在执行层做校验
