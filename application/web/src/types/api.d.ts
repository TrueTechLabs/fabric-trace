// API通用类型定义

/**
 * HTTP方法类型
 */
export type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

/**
 * Axios配置类型
 */
export interface AxiosConfig {
  url: string
  method: HttpMethod
  data?: any
  params?: any
  headers?: Record<string, string>
  timeout?: number
}

/**
 * API响应基础类型
 */
export interface BaseResponse<T = any> {
  code: number
  message: string
  data: T
}

/**
 * 分页请求参数
 */
export interface PageParams {
  page: number
  pageSize: number
}

/**
 * 分页响应数据
 */
export interface PageResponse<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}
