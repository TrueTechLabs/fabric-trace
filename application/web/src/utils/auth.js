import Cookies from 'js-cookie'

const TokenKey = 'vue_admin_template_token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token, {
    expires: 7, // 7天过期
    secure: process.env.NODE_ENV === 'production', // 生产环境使用HTTPS
    sameSite: 'strict' // 防止CSRF攻击
  })
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}
