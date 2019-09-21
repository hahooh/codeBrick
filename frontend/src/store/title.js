export default {
    state: {
        title: ''
    },

    mutations: {
        setTitle(state, { newTitle }) {
            state.title = newTitle
        }
    },

    actions: {

    },

    getters: {
        getTitle(state) {
            return state.title;
        }
    }
}