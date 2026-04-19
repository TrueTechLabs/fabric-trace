import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    if (store.getters.token) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['Authorization'] = getToken()
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    // console.log(res)
    // if the custom code is not 200, it is judged as an error.
    if (res.code !== 200) {
      Message({
        message: res.message || 'Error',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      return res
    }
  },
  error => {
    let message = '操作失败'
    if (error.response) {
      switch (error.response.status) {
        case 401:
          message = '登录已过期，请重新登录'
          // 401自动跳转登录
          store.dispatch('user/resetToken').then(() => {
            location.href = '/login'
          })
          break
        case 403:
          message = '没有权限执行此操作'
          break
        case 404:
          message = '请求的资源不存在'
          break
        case 500:
          message = '服务器错误，请稍后重试'
          break
        case 502:
        case 503:
        case 504:
          message = '服务暂时不可用，请稍后重试'
          break
        default:
          message = error.response.data?.message || error.response.data?.error || '请求失败'
      }
    } else if (error.message.includes('timeout')) {
      message = '请求超时，请检查网络连接'
    } else if (error.message.includes('Network')) {
      message = '网络连接失败，请检查网络'
    } else if (error.message) {
      message = error.message
    }
    Message({
      message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
