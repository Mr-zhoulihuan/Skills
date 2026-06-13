# Skills — Codex AI 技能集

Codex AI 的技能集合仓库。每个技能封装了特定场景下的指令和工作流程，让 Codex 在对应任务中表现更精准、高效。

## 技能列表

| 技能 | 说明 |
|------|------|
| [token-saver](./token-saver/SKILL.md) | 在对话和编码中减少 token 消耗，通过精简回复、高效使用工具、避免冗余上下文加载等方式提升 token 利用率 |

## 安装方式

### 从 GitHub 安装

在 Codex 中直接使用 skill-installer：

```text
安装 Skills 仓库的 token-saver
```

### 手动安装

1. 将对应技能目录（如 `token-saver`）复制到 Codex 的 skills 目录：
   - Codex Desktop: `$CODEX_HOME/skills/`
   - Codex CLI: `~/.codex/skills/`
2. 重新加载 Codex 即可生效。

## 使用方法

技能安装后，在相应场景下 Codex 会自动触发对应的指令集。你也可以在对话中明确引用技能名称来启用它，例如：

- `$token-saver` — 启用 token 节省模式

## 贡献

欢迎提交新的技能或改进现有技能。每个技能应包含：

- `SKILL.md` — 核心指令和规则
- `agents/` (可选) — 界面展示用的元数据和默认提示词
- `references/` (可选) — 参考文档和变体说明
- `scripts/` (可选) — 辅助脚本
- `assets/` (可选) — 图片、模板等静态资源

提交前请确保技能描述清晰，触发规则明确，且与现有技能不冲突。

## 仓库结构

```
Skills/
├── README.md
├── token-saver/
│   ├── SKILL.md
│   └── agents/
│       └── openai.yaml
```
