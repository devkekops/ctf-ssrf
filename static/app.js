var app = new Vue({
    el: '#app',
    data: {
        url: '',
        result: '',
        visibility: false
    },
    methods: {
        convert: function() {
            this.visibility = false
            this.result = '';
            axios
                .get('http://0.0.0.0/convert?url=' + this.url)
                .then(response => (
                    console.log(response.data),
                    this.visibility = true,
                    this.result = response.data));
        },
    },
})