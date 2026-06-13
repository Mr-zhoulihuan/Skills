 # AI Agent 学习资源汇总
 
 ---
 
 ## 必读论文
 
 | 论文 | 年份 | 核心贡献 | 链接 |
 |---|---|---|---|
 | ReAct: Synergizing Reasoning and Acting in Language Models | 2022 | 提出 ReAct 范式，Agent 基础 | https://arxiv.org/abs/2210.03629 |
 | Toolformer: Language Models Can Teach Themselves to Use Tools | 2023 | LLM 自学习工具调用 | https://arxiv.org/abs/2302.04761 |
 | Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks | 2020 | RAG 开山之作 | https://arxiv.org/abs/2005.11401 |
 | Tree of Thoughts: Deliberate Problem Solving | 2023 | LLM 树状推理 | https://arxiv.org/abs/2305.10601 |
 | Generative Agents: Interactive Simulacra of Human Behavior | 2023 | 长期记忆 + Agent 架构 | https://arxiv.org/abs/2304.03442 |
 
 ---
 
 ## 框架与工具
 
 | 工具 | 用途 | 链接 |
 |---|---|---|
 | LangChain / LangGraph | 通用 Agent 框架 | https://github.com/langchain-ai/langchain |
 | CrewAI | 多 Agent 协作框架 | https://github.com/joaomdmoura/crewAI |
 | OpenAI Agents SDK | OpenAI 官方 Agent SDK | https://github.com/openai/openai-agents-python |
 | LlamaIndex | RAG 框架 | https://github.com/run-llama/llama_index |
 | AutoGen | 微软多 Agent 框架 | https://github.com/microsoft/autogen |
 | Semantic Kernel | 微软 Agent SDK（.NET/Python） | https://github.com/microsoft/semantic-kernel |
 | Dify | 可视化 Agent / RAG 平台 | https://github.com/langgenius/dify |
 | Coze | 字节跳动的 Agent 搭建平台 | https://www.coze.com |
 
 ---
 
 ## 上手代码仓库
 
 | 仓库 | 说明 | 链接 |
 |---|---|---|
 | openai-cookbook | OpenAI 官方代码示例 | https://github.com/openai/openai-cookbook |
 | awesome-llm-apps | LLM 应用合集 | https://github.com/Shubhamsaboo/awesome-llm-apps |
 | ragflow | 开源 RAG 引擎 | https://github.com/infiniflow/ragflow |
 | langgraph-examples | LangGraph 官方示例 | https://github.com/langchain-ai/langgraph-examples |
 
 ---
 
 ## 课程与教程
 
 | 课程 | 平台 | 链接 |
 |---|---|---|
 | Hugging Face Agents Course | 免费在线 | https://huggingface.co/learn/agents-course/ |
 | DeepLearning.AI — Building Agentic RAG | 付费（短课） | https://www.deeplearning.ai/short-courses/ |
 | LangChain 官方教程 | 免费在线 | https://python.langchain.com/docs/tutorials/ |
 | 吴恩达 — Function Calling 与 Tool Use | 免费 | https://www.deeplearning.ai/short-courses/ |
 | Stanford CS224N (NLP) | 免费大学课 | https://web.stanford.edu/class/cs224n/ |
 
 ---
 
 ## 书籍
 
 | 书名 | 说明 |
 |---|---|
 | 《Building LLM Apps》(O'Reilly) | LLM 应用构建实战，含 Agent |
 | 《AI Engineering》(O'Reilly) | Chip Huyen 著，涵盖 RAG / Agent / 评估 |
 | 《大语言模型实战指南》 | 中文，偏向工程落地 |
 
 ---
 
 ## 实践建议
 
 ### 环境准备
 
 ```bash
 pip install openai chromadb langchain crewai pypdf
 ```
 
 ### 推荐练手项目（按难度排序）
 
 1. **个人知识库问答 Bot** — 喂几篇 PDF，能问能答（RAG 入门）
 2. **自动化调研助手** — 给一个主题，自动搜索、整理、生成报告（Agent 入门）
 3. **多 Agent 内容工厂** — 写稿 Agent + 配图 Agent + 审核 Agent + 发布 Agent（综合）
 4. **Slack / 飞书 Bot** — 对接即时通讯，Agent 自动回复（生产级）
 
 ### 好习惯
 
 - 每次实验记录：模型、参数、Prompt、结果、失败原因
 - 先手写再框架，理解原理后再使用工具的封装
 - Agent 的代码要有完整的日志：记录每一轮 Thought / Action / Observation
 - 把每次工具的调用耗时、token 消耗记录下来，优化成本
 
 ---
 
 *祝你学得扎实，代码不崩。*
