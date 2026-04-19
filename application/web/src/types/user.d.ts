// 用户相关类型定义

/**
 * 用户角色类型
 */
export type UserRole = '种植户' | '工厂' | '运输司机' | '商店' | '消费者'

/**
 * 用户信息
 */
export interface UserInfo {
  name: string
  userType: UserRole
  avatar?: string
  roles?: string[]
}

/**
 * 登录表单
 */
export interface LoginForm {
  username: string
  password: string
}

/**
 * 注册表单
 */
export interface RegisterForm {
  username: string
  password: string
  password2: string
  userType: UserRole
}

/**
 * Token响应
 */
export interface TokenResponse {
  token: string
  txid?: string
}
