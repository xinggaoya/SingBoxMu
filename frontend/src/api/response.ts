import ky from 'ky';

// 创建一个ky实例，可以设置默认配置
const apiClient = ky.create({
    prefixUrl: 'http://127.0.0.1:20123', // 设置默认API URL前缀
    headers: {
        'Authorization': 'Bearer CDBEA571-B1FC-4AEF-A8BE-1E5B3D11B9C2', // 设置默认请求头
        'Content-Type': 'application/json',
    },
    timeout: 5000, // 设置默认超时时间
});

// 封装GET请求
export const get = async (url: string, params: any, onProgress?: (progress: any, chunk: Uint8Array) => void) => {
    const response = await apiClient.get(url, {
        searchParams: params,
        onDownloadProgress: onProgress,
    });
    return response.json();

};

// 封装POST请求
export const post = async (url: string, data: any, onProgress?: () => void) => {
    const response = await apiClient.post(url, {
        json: data,
        onDownloadProgress: onProgress,
    });
    return response.json();
};

// 封装PUT请求
export const put = async (url: string, data: any) => {
    const response = await apiClient.put(url, {
        json: data,
    });
    if (response.status === 204) {
        return;
    }
    return response.json();

};

// 封装DELETE请求
export const del = async (url: string) => {
    const response = await apiClient.delete(url);
    return response.json();

};
