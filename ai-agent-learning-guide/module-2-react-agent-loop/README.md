 # 模块二：Agent 循环与推理（ReAct / Agent Loop）
 
 ## 是什么
 
 **Agent Loop（Agent 循环）** 是 AI Agent 的核心引擎。它让 LLM 不再是一次回答一个问题，而是在一个循环中持续：**思考 → 行动 → 观察结果 → 再思考**，直到任务完成。
 
 最经典的范式是 **ReAct（Reason + Act）**，由 Shunyu Yao 等人在 2022 年提出。它把推理链（Chain-of-Thought）和工具调用结合在一起。
 
 ```
 循环开始
   ↓
 ① Thought（思考）—— "我需要先查北京的天气"
   ↓
 ② Action（行动）—— 调用 get_weather(city="北京")
   ↓
 ③ Observation（观察）—— "北京 25°C 晴"
   ↓
 ④ Thought（再思考）—— "那再查一下上海的"
   ↓
 ⑤ Action（行动）—— 调用 get_weather(city="上海")
   ↓
 ⑥ Observation（观察）—— "上海 30°C 多云"
   ↓
 ⑦ Thought（再思考）—— "上海比北京暖，可以给用户对比答案了"
   ↓
 ⑧ Final Answer（最终回答）—— 输出给用户
 循环结束
 ```
 
 ---
 
 ## 核心概念
 
 - **ReAct**：Reason + Act 的缩写，推理和行动交替进行
 - **Agent Loop**：管理循环的执行引擎，包括最大步数限制、停止条件
 - **Planning**：Agent 在行动前分解任务的能力（如 Plan-and-Execute）
 - **Reflection**：Agent 对自身行动结果的自我评估和修正
 - **Multi-Agent**：多个 Agent 分工协作（一个写代码、一个审查、一个执行）
 
 ---
 
 ## 学习路径
 
 ### 第 1 步：手写一个 ReAct 循环（一天 — 最重要）
 
 用纯 Python 实现一个最小 ReAct Agent，**不要用任何框架**。
 
 ```python
 import json
 
 MAX_ITERATIONS = 10
 
 def agent_loop(user_input):
     messages = [{"role": "user", "content": user_input}]
 
     for i in range(MAX_ITERATIONS):
         response = llm(messages)  # 调用 LLM
         content = response.choices[0].message.content
 
         if "Final Answer:" in content:
             return content.split("Final Answer:")[1].strip()
 
         # 解析 Action
         action = parse_action(content)  # 从文本中提取 Action 行
         result = execute_tool(action["name"], action["args"])
 
         messages.append({"role": "assistant", "content": content})
         messages.append({"role": "user", "content": f"Observation: {result}"})
 
     return "Max iterations reached."
 ```
 
 这个过程能帮你深刻理解：
 - Agent 不是在"思考"，是在**根据格式生成文本**
 - 循环控制、停止条件、历史维护都是你的代码在管
 - 失败场景（循环死锁、越界）怎么兜底
 
 ### 第 2 步：增加规划能力（半天）
 
 - 实现 Plan-and-Execute：Agent 先输出一个任务分解计划，再逐条执行
 - 实现子任务合并：如果一个子任务失败，Agent 能否重新规划
 
 ```python
 # Plan 格式示例
 Plan:
 1. 搜索"2024 年 AI 市场规模"
 2. 搜索"2025 年 AI 市场预测"
 3. 对比两组数据，生成报告
 ```
 
 ### 第 3 步：增加错误恢复与重试（半天）
 
 - 工具调用超时 → Agent 重试或换方案
 - 工具返回空结果 → Agent 换关键词搜索
 - Agent 陷入死循环 → 人为中断并回退
 
 ### 第 4 步：Multi-Agent 协作（一天）
 
 - 实现 2-3 个 Agent：编写、审查、执行
 - 探索不同的通信方式：共享 Message Queue / 共享上下文 / 函数调用
 
 ### 第 5 步：用框架实现（半天）
 
 - 用 LangChain AgentExecutor 或 CrewAI 重写
 - 对比框架与手写的差异
 
 ---
 
 ## 资源推荐
 
 | 资源 | 类型 | 链接 |
 |---|---|---|
 | ReAct 原始论文 | 论文 | https://arxiv.org/abs/2210.03629 |
 | LangChain Agent 概念 | 文档 | https://python.langchain.com/docs/concepts/agents/ |
 | OpenAI Agents SDK | 代码 | https://github.com/openai/openai-agents-python |
 | Anthropic Agent 设计指南 | 文档 | https://docs.anthropic.com/en/docs/build-with-claude/agentic |
 | CrewAI 官方教程 | 文档 | https://docs.crewai.com/ |
 | Hugging Face Agent 课程 | 教程 | https://huggingface.co/learn/agents-course/ |
 
 ---
 
 ## 检验标准
 
 - [ ] 能不依赖框架手写一个可运行的 ReAct Agent
 - [ ] 能处理 Agent 陷入死循环（最大步数 / 重复检测）
 - [ ] 能实现 Plan-and-Execute 模式
 - [ ] 能说清楚 Multi-Agent 的通信和协调机制
 - [ ] 能用 CrewAI / LangGraph 搭一个简单的多 Agent 系统
 
 ---
 
 ## 常见误区
 
 - ❌ **过早引入框架** — LangChain 封装的 AgentExecutor 很方便，但也把核心循环遮住了。先手写再框架
 - ❌ **忽略状态管理** — Agent 把每轮结果全塞给 LLM，不知道 truncate / summarize history，token 很快爆炸
 - ❌ **没有停止条件** — Agent 会一直循环下去，必须有明确的终止逻辑
 - ❌ **把 Multi-Agent 当成"更能干"的 Agent** — 多 Agent 解决的是角色分离，不是性能提升
 - ❌ **期望 Agent 每次都正确** — Agent 本质是概率性的，系统设计要考虑降级和人工介入
