// 模拟 API 函数，所有调用都返回提示
export function uploadLoginBackground() {
  return new Promise((resolve) => {
    resolve({
      code: 200,
      data: { url: 'https://via.placeholder.com/800x600' },
      message: '演示模式：购买课程后可用'
    })
  })
}

export function saveConfig() {
  return new Promise((resolve) => {
    resolve({
      code: 200,
      message: '演示模式：购买课程后可用'
    })
  })
}
