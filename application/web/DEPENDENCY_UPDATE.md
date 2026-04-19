# 依赖更新说明

## 当前过时依赖

以下依赖版本较旧，建议更新：

### 核心依赖
- `axios`: 0.18.1 → 1.x (重大版本更新)
- `element-ui`: 2.13.2 → 2.15.x (最新版本)
- `vue`: 2.6.10 → 2.7.x (Vue 2最新版本)

### 开发依赖
- `@vue/cli-service`: 4.4.4 → 5.x (需要迁移)
- `sass`: 1.26.8 → 1.x (更新)
- `eslint`: 6.7.2 → 8.x (需要配置迁移)

## 更新步骤

### 方案1：安全更新（推荐）
仅更新补丁版本，最小化风险：

```bash
# 进入前端目录
cd application/web

# 更新补丁版本
npm update

# 审计依赖漏洞
npm audit fix
```

### 方案2：完整更新（需要测试）
更新到最新版本：

```bash
# 更新核心依赖
npm install axios@latest element-ui@latest

# 更新Vue到2.7版本
npm install vue@2.7.x vue-template-compiler@2.7.x
```

## 注意事项

1. **axios 0.x → 1.x 破坏性更改**：
   - 需要检查所有API调用代码
   - 某些配置选项可能需要调整

2. **Vue CLI 4 → 5**：
   - 需要运行 `vue upgrade`
   - webpack配置可能需要调整
   - node-sass已被弃用，使用sass

3. **测试建议**：
   - 更新前先备份代码
   - 在开发环境充分测试
   - 检查控制台是否有警告或错误
   - 测试所有核心功能

## 当前状态

由于依赖更新可能影响项目稳定性，建议：
1. 先完成功能测试
2. 在稳定后再进行依赖更新
3. 使用 `npm outdated` 查看过时依赖

## 更新命令参考

```bash
# 查看过时依赖
npm outdated

# 安全更新
npm update

# 修复安全漏洞
npm audit fix

# 强制修复（可能破坏性更改）
npm audit fix --force
```
