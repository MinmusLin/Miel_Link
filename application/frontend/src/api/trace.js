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

export function getProductHistory(data) {
    return request({
        url: '/getProductHistory',
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data
    })
}