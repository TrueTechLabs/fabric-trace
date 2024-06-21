import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

export function register(data) {
  return request({
    url: '/register',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/getInfo',
    method: 'post',
    headers: {
      'Authorization': `Bearer ${token}`
    }
    // params: { token }
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}
