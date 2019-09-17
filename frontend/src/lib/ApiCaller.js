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

    execute(method, path, params) {
        return axios({
            method,
            url: path,
            params,
            baseURL: this.baseUrl
        })
    }
}