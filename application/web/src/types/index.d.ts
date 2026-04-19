// 类型定义导出
export * from './trace'
export * from './user'
export * from './api'

// Vue组件扩展类型
declare module 'vue/types/vue' {
  interface Vue {
    // 可以在这里添加全局属性类型
  }
}
