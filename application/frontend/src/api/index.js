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
    })
}

export function logout() {
    return request({
        url: '/logout',
        method: 'post'
    })
}

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

export function ipfsDownload(data) {
    return request({
        url: '/ipfsDownload',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}

export function ipfsUpload(data) {
    return request({
        url: '/ipfsUpload',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}

export function getProductInfo(data) {
    return request({
        url: '/getProductInfo',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}

export function getProductList(data) {
    return request({
        url: '/getProductList',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}

export function getAllProductInfo(data) {
    return request({
        url: '/getAllProductInfo',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}
