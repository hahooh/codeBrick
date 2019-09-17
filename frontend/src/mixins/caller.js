import ApiCaller from '../lib/ApiCaller'

export const caller = {
    methods: {
        getApiCaller() {
            return new ApiCaller('http://localhost:4001')
        }
    }
}