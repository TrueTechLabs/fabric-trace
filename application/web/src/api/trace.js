import request from '@/utils/request'

/**
 * 数据上链（需要文件上传，使用FormData）
 */
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

/**
 * 获取单个农产品信息（使用FormData格式以兼容后端PostForm）
 */
export function getFruitInfo(traceabilityCode) {
  const formData = new FormData()
  formData.append('traceability_code', traceabilityCode)
  return request({
    url: '/getFruitInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data: formData
  })
}

/**
 * 获取农产品列表（使用JSON格式）
 */
export function getFruitList() {
  return request({
    url: '/getFruitList',
    method: 'post'
  })
}

/**
 * 获取所有农产品信息（使用JSON格式）
 */
export function getAllFruitInfo() {
  return request({
    url: '/getAllFruitInfo',
    method: 'post'
  })
}

/**
 * 获取农产品历史记录（使用FormData格式以兼容后端PostForm）
 */
export function getFruitHistory(traceabilityCode) {
  const formData = new FormData()
  formData.append('traceability_code', traceabilityCode)
  return request({
    url: '/getFruitHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data: formData
  })
}

