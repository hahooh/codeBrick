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

        updateItem(contenxt, { path, id, updates }) {
            return apiCaller.put(`${path}/${id}`, updates)
                .then(response => {
                    const updatedItem = response.data;
                    const newItemList = this.items.filter(item => item.Id !== updatedItem.Id);
                    newItemList.push(updatedItem);
                    contenxt.saveItems({ items: newItemList });
                })
                .catch(error => {
                    console.log(error);
                });
        },

        createItem(context, { path, item }) {
            return apiCaller.post(path, item)
                .then(response => {
                    const newItemList = this.items.concat([response.data]);
                    context.saveItems({ items: newItemList });
                })
                .catch(error => {
                    console.log(error);
                });
        },

        delete(context, { path, id }) {
            return apiCaller.delete(`${path}/${id}`)
                .then(response => {
                    const newItemList = this.items.filter(item => item.Id != id);
                    context.saveItems({ items: newItemList })
                })
                .catch(error => {
                    console.log(error);
                });
        }
    },

    getters: {
        
    }
}