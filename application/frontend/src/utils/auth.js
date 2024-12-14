import Cookies from 'js-cookie'

const TokenKey = 'vue_admin_template_token'

export function getToken() {
    // noinspection JSUnresolvedReference
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    // noinspection JSUnresolvedReference
    return Cookies.set(TokenKey, token)
}

export function removeToken() {
    // noinspection JSUnresolvedReference
    return Cookies.remove(TokenKey)
}
