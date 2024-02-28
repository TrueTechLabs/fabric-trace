import request from '@/utils/request'

export function uplink(data) {
  return request({
    url: '/uplink',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getFruitInfo
export function getFruitInfo(data) {
  return request({
    url: '/getFruitInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getFruitList
export function getFruitList(data) {
  return request({
    url: '/getFruitList',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getAllFruitInfo
export function getAllFruitInfo(data) {
  return request({
    url: '/getAllFruitInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getFruitHistory
export function getFruitHistory(data) {
  return request({
    url: '/getFruitHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

