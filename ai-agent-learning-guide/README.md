# AI Agent 学习技能

 AI Agent 是当下 LLM 应用开发最核心的方向，本质是让语言模型具备感知、推理、调用工具、记忆和行动的能力。在动手搭建 Agent 之前，先深入理解 LLM 本身的工作原理，会让后续学习事半功倍。

 本指南在 Agent 三大核心模块之前，增加 **LLM 基础理解** 模块，帮助你在掌握工具之前先理解模型，每模块独立成章，包含概念介绍、学习路径和资源推荐。

---

 ## 四大模块总览

 | 模块 | 核心能力 | 一句话概括 |
 |---|---|---|
 | [模块一：工具调用](/module-1-function-calling/README.md) | LLM 输出结构化参数 → 调用外部函数/API | Agent 的"手" |
 | [模块二：Agent 循环与推理](/module-2-react-agent-loop/README.md) | 思考 → 行动 → 观察 → 再思考的循环 | Agent 的"大脑" |
 | [模块三：RAG 与记忆系统](/module-3-rag-memory/README.md) | 从外部知识库检索信息 + 多轮对话记忆 | Agent 的"长期记忆" |
 | [模块四：深入理解 LLM](/module-4-llm-understanding/README.md) | Token、Attention、Transformer、预训练、推理参数 | 给你"翻译"大模型的黑盒 |

 ---

 ## 推荐学习顺序
 
```
 第 1 步 ── 模块四：深入理解 LLM（选读，推荐先看）
          ↓ 理解 LLM 的本质：预测下一个 Token
 第 2 步 ── 模块一：工具调用（Function Calling）
          ↓ 理解 LLM 如何"调用"外部能力
 第 3 步 ── 模块二：Agent 循环（ReAct）
          ↓ 理解 Agent 如何自主推理和决策
 第 4 步 ── 模块三：RAG & 记忆
          ↓ 理解知识注入与历史对话管理
          ↓
    最终 ── 三者组合，搭建完整 Agent
 ```
 
 ---
 
 ## 前置要求
 
 - Python 3.9+，熟悉基础语法
 - 了解 HTTP API 的基本概念
 - 有 OpenAI / 国产大模型的 API Key
 - 一台能跑代码的机器（普通笔记本即可）
 
 ---
 
 ## 各模块快速入口
 
 - [模块一：工具调用 →](/module-1-function-calling/README.md)
 - [模块二：Agent 循环与推理 →](/module-2-react-agent-loop/README.md)
 - [模块三：RAG 与记忆系统 →](/module-3-rag-memory/README.md)
 - [模块四：深入理解 LLM →](/module-4-llm-understanding/README.md)
 - [资源汇总 →](/resources.md)
