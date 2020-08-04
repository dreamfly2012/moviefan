import axios from './api'

export function getMovie (params = {}) { // 从外部接受参数，没有参数默认为空对象
    console.log(params)
    console.log('sss')
    return axios.get('/get', params) // return对应的get/post方法，第一个填路径，第二个给参数对象
}