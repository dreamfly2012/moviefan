import axios from 'axios'
axios.defaults.baseURL = 'http://127.0.0.1:8080'
axios.defaults.withCredentials = false
axios.defaults.timeout = 100000
// // axios拦截器
axios.interceptors.request.use(
   config => {
      let token = sessionStorage.getItem('access_token')
   
      if (token) {
         config.headers = {
         'access-token': token,
         'Content-Type': 'application/x-www-form-urlencoded'
         }
      }
      if (config.url === 'refresh') {
         config.headers = {
         'refresh-token': sessionStorage.getItem('refresh_token'),
         'Content-Type': 'application/x-www-form-urlencoded'
         }
      }
      return config
   }
)
 
axios.interceptors.response.use(response => {
   
     // 在这里你可以判断后台返回数据携带的请求码
    if (response.status === 200 ) {
      return response.data.data || response.data
    }else {
      // 非200请求抱错
      throw Error(response.data.msg || '服务异常')
    }
 })
export default axios