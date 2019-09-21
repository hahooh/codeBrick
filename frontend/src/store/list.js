import { caller } from '../mixins/caller'
const apiCaller = caller.methods.getApiCaller();

export default {
    state: {
        items: [],
        item: null
    },

    mutations: {
        saveItems: (state, { items }) => {
            state.items = items
        },

        saveItem: (state, { item }) => {
            state.item = item
        },

        deleteItem: (state, { itemId }) => {
            state.items = state.items.filter(item => item.Id !== itemId);
        },

        addItem: (state, { item }) => {
            state.items.push(item);
        }
    },

    actions: {
        fetchItems(context, { path }) {
            return apiCaller.get(path)
                .then(response => {
                    context.commit('saveItems', { items: response.data });
                    return Promise.resolve(response);
                })
                .catch(error => {
                    return Promise.reject(error);
                });
        },

        fetchItem(context, { path, id }) {
            return apiCaller.get(`${path}/${id}`)
                .then(response => {
                    context.commit('saveItem', { item: response.data });
                    return Promise.resolve(response);
                })
                .catch(error => {
                    console.log(error);
                    return Promise.reject(error);
                });
        },

        updateItem(context, { path, id, updates }) {
            return apiCaller.put(`${path}/${id}`, updates)
                .then(response => {
                    const updatedItem = response.data;
                    context.commit('deleteItem', { itemId: updatedItem.Id });
                    context.commit('addItem', { item: updatedItem });
                })
                .catch(error => {
                    console.error(error);
                });
        },

        createItem(context, { path, item }) {
            return apiCaller.post(path, item)
                .then(response => {
                    context.commit('addItem', { item: response.data });
                })
                .catch(error => {
                    console.error(error);
                });
        },

        delete(context, { path, id }) {
            return apiCaller.delete(`${path}/${id}`)
                .then(response => {
                    context.commit('addItem', { itemId: id });
                })
                .catch(error => {
                    console.error(error);
                });
        }
    },

    getters: {

    }
}