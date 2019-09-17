import { subtitles } from '../router'

export const url = {
    data: function () {
        return {
            subtitles,
        }
    },

    methods: {
        getTitleByUrl: function (url) {
            if (url === '/') {
                return 'Home'
            }
            const subtitle = subtitles.find(subtitle => subtitle.url === url);
            return subtitle ? subtitle.title : 'Page Not Found';
        },
        getHeadersByUrl(url) {
            if (url === '/') {
                return [];
            }
            const subtitle = subtitles.find(subtitle => subtitle.url === url);
            if (subtitle && subtitle.headers) {
                return subtitle.headers
            }
            return [];
        }
    }
}