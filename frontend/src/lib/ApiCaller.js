import axios from 'axios'

export default class ApiCaller {
    constructor(baseUrl) {
        this.baseUrl = baseUrl;
    }

    get(path, params) {
        return this.execute('get', path, params);
    }

    post(path, params) {
        return this.execute('post', path, params);
    }

    put(path, params) {
        return this.execute('put', path, params);
    }

    delete(path, params) {
        return this.execute('delete', path, params);
    }

    execute(method, path, data) {
        const requestOptions = {
            method,
            url: path,
            baseURL: this.baseUrl,
        }

        if(method === 'get') {
            requestOptions.params = data;
        } else {
            requestOptions.data = data;
        }

        return axios(requestOptions)
    }
}