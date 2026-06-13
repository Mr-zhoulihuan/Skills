 # 模块三：RAG 与记忆系统（Retrieval-Augmented Generation & Memory）
 
 ## 是什么
 
 RAG（检索增强生成）和记忆系统解决同一个问题：**LLM 不知道的事情怎么办？**
 
 - **RAG**：从外部知识库（文档库、数据库、搜索引擎）检索相关信息，注入到 LLM 上下文中，让 LLM 基于这些信息回答。解决的是"知识不足"的问题。
 - **记忆**：管理多轮对话中的历史信息，让 Agent 能记住之前的对话。解决的是"上下文丢失"的问题。
 
 ```
 用户提问 ──→ 检索（从知识库找到相关内容）
               ↓
           [检索结果 + 对话历史 + 用户提问] 拼成 Prompt
               ↓
           LLM 基于上下文生成回答
               ↓
           → 回答给用户 + 更新对话记忆
 ```
 
 ---
 
 ## 核心概念
 
 - **Embedding**：将文本转为向量（浮点数数组），用于语义匹配
 - **向量数据库**：存储向量并提供相似度检索（余弦距离 / 欧氏距离）
 - **Chunking（分块）**：将长文档切成小块，保证检索质量
 - **Retrieval 策略**：关键词搜索、向量搜索、混合搜索（Hybrid Search）
 - **Reranking（重排序）**：对检索结果重新排序，提高精度
 - **记忆类型**：短期记忆（窗口）、长期记忆（摘要）、持久化记忆（向量库）
 
 ---
 
 ## 学习路径
 
 ### 第 1 步：跑通 RAG 最小链路（半天）
 
 ```python
 # 最小链路：文档 → 分块 → Embedding → 存储 → 检索 → 生成
 
 from openai import OpenAI
 
 # 1. 读文档 → 分块
 chunks = split_text(document, chunk_size=500, overlap=50)
 
 # 2. 生成 Embedding
 embeddings = client.embeddings.create(
     model="text-embedding-3-small",
     input=chunks
 )
 
 # 3. 存入内存/向量库（简单用列表模拟）
 vector_store = [(emb.embedding, chunk) for emb, chunk in zip(embeddings.data, chunks)]
 
 # 4. 用户提问 → 检索
 query_emb = client.embeddings.create(
     model="text-embedding-3-small", input=question
 ).data[0].embedding
 
 results = search_by_similarity(vector_store, query_emb, top_k=3)
 
 # 5. 拼接上下文 → LLM 回答
 response = client.chat.completions.create(
     model="gpt-4o",
     messages=[
         {"role": "system", "content": f"基于以下资料回答：\n{results}"},
         {"role": "user", "content": question}
     ])
 ```
 
 ### 第 2 步：引入向量数据库（半天）
 
 选一个：
 - **Chroma**（本地，最简单）
 - **Qdrant**（本地或 Docker，性能好）
 - **Milvus**（生产级，部署复杂些）
 
 练习：用 Chroma 替代上面的内存存储。
 
 ### 第 3 步：优化检索质量（一天）
 
 - 对比不同分块策略（chunk_size 300/500/1000，overlap 0/50/100）
 - 加入 Reranker 提升精度（`Cohere Rerank` 或 `BGE Reranker`）
 - 实现 Hybrid Search（BM25 关键词 + 向量检索，结果融合）
 - 处理"检索不到"的情况：降级回答、反问用户、换检索源
 
 ### 第 4 步：实现记忆系统（一天）
 
 - **短期记忆**：维护一个滑动窗口（最近 N 轮对话）
 - **摘要记忆**：当对话超长时，对历史做摘要压缩
 - **向量记忆**：把重要的对话存入向量库，在需要时检索
 
 ```python
 # 记忆接口示例
 class Memory:
     def add(self, role, content):        # 添加一条对话
     def get_recent(self, n=10):           # 获取最近 N 条
     def summarize(self):                  # 压缩历史为摘要
     def search(self, query, k=3):         # 检索相关历史
 ```
 
 ### 第 5 步：Agent + RAG + 记忆 三者整合（一天）
 
 把模块一（工具调用）、模块二（Agent 循环）、模块三（RAG + 记忆）组合成一个完整的 Agent：
 
 - Agent 可以搜索知识库 → 分析 → 调用工具 → 记住用户偏好
 - 完整链路：用户问 → 检索知识库 → Agent 思考 → 调用 API → 记住结论 → 回答
 
 ---
 
 ## 资源推荐
 
 | 资源 | 类型 | 链接 |
 |---|---|---|
 | LangChain RAG 教程 | 文档 | https://python.langchain.com/docs/tutorials/rag/ |
 | Chroma 向量库入门 | 文档 | https://docs.trychroma.com/ |
 | Pinecone RAG 指南 | 教程 | https://www.pinecone.io/learn/rag/ |
 | Cohere Rerank API | 文档 | https://docs.cohere.com/reference/rerank |
 | BM25 算法（rank_bm25） | 库 | https://pypi.org/project/rank-bm25/ |
 | LlamaIndex RAG 框架 | 文档 | https://docs.llamaindex.ai/ |
 | 阿里云 RAG 最佳实践 | 文档 | https://help.aliyun.com/zh/model-studio/rag |
 | OpenAI Embedding 文档 | 文档 | https://platform.openai.com/docs/guides/embeddings |
 
 ---
 
 ## 检验标准
 
 - [ ] 能不依赖框架跑通完整 RAG 链路（文档→分块→向量化→检索→生成）
 - [ ] 能说清楚 Chunk Size、Overlap、Top-K 对结果的影响
 - [ ] 能实现 Hybrid Search（向量 + 关键词），并理解两者的适用场景
 - [ ] 能实现至少两种记忆策略（窗口 + 摘要 / 窗口 + 向量）
 - [ ] 能把 RAG + 记忆整合到 Agent 中
 
 ---
 
 ## 常见误区
 
 - ❌ **认为"向量检索解决一切"** — 关键词+向量混合搜索在很多场景比纯向量更好
 - ❌ **分块太粗暴** — 分块策略直接影响检索质量，不同的文档结构（Markdown/PDF/代码）需要不同的 Chunking 策略
 - ❌ **不处理检索空洞** — 检索不到相关内容时，LLM 会强行编造，必须有降级策略
 - ❌ **记忆无限增长** — 每轮对话全塞进 Prompt，token 爆炸。必须做截断或摘要
 - ❌ **Embedding 模型选不对** — 中文场景用 `text-embedding-3-small`（多语言支持好）或国产 Embedding 模型，效果好于纯英文模型
 - ❌ **忽视 Reranker** — Top-K 检索结果里可能有噪音，Reranker 能显著提升最终回答质量
