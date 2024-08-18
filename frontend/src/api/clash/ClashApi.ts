import {get, put} from '@/api/response'


// 获取日志
export const GetLogs = (onProgress: (progress: any, chunk: Uint8Array) => void) => {
    return get('logs', {}, onProgress)
}
// 获取流量
export const GetTraffic = (onProgress: (progress: any, chunk: Uint8Array) => void) => {
    return get('traffic', {}, onProgress)
}
// 切换代理
export const SwitchProxy = (group: string, name: string) => {
    return put('proxies/' + group, {
        name
    })
}
// 获取内存
export const GetMemory = (onProgress: (progress: any, chunk: Uint8Array) => void) => {
    return get('memory', {}, onProgress)
}

// 获取版本
export const GetVersion = () => {
    return get('version', {})
}

// 获取代理
export const GetProxies = () => {
    return get('proxies', {})
}