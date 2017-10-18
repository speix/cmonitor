Vue.component('modal', {
    template: `

            <div class="modal is-active">
                <div class="modal-background"></div>
                <div class="modal-card">
                    <header class="modal-card-head">
                        <p class="modal-card-title">
                            <slot name="modal-title"></slot>
                        </p>
                        <button class="delete" aria-label="close" @click="$emit('close')"></button>
                    </header>
                    <section class="modal-card-body">
                        <slot name="modal-content"></slot>
                    </section>
                    <footer class="modal-card-foot">
                        <button class="button" @click="$emit('close')">Close</button>
                    </footer>
                </div>
            </div>

            `
});

var vm = new Vue({
    el: '#root',
    data: {
        list: '',
        lists: [],
        subscribers: [],
        showSubscribersModal: false,
        showAddSubscriberForm: false,
        modalTitle: '',
        no_records_message: '',
        sEmail: '',
        sName: ''
    },
    created: function(){
        this.getCampaignLists();
    },
    methods: {
        getCampaignLists: function(){
            this.$http.get("lists").then(function(response) {
                var self = this;

                response.json().then(function(data) {
                    self.lists = data;
                });

            }).catch(function(response){
                alert(response.body.message);
            });
        },
        GetListSubscribers: function(list, list_name, showAddSubscriberForm) {
            this.showSubscribersModal = true;
            this.showAddSubscriberForm = showAddSubscriberForm;
            this.modalTitle = list_name;
            this.subscribers = [];
            this.list = '';
            this.no_records_message = '';

            this.$http.get("subscribers?list=" + list).then(function(response) {
                var self = this;

                response.json().then(function(data) {
                    self.subscribers = data.subscribers;
                    self.list = data.listid;

                    if (data.subscribers.length == 0){
                        self.no_records_message = 'No records found in ' +list_name+ ' List';
                    }
                });

            }).catch(function(response){
                alert(response.body.message);
            });
        },
        DeleteSubscriber: function(list_id, list_name, email, index) {

            this.$http.delete("subscriber/delete", {params: {email: email, list: list_id}}).then(function(response) {
                var self = this;

                response.json().then(function(data){

                    vm.subscribers.splice(index, 1);

                    if (self.subscribers.length === 0){
                        self.no_records_message = 'No records found in ' +list_name+ ' List';
                    }

                });

            }).catch(function(response){
                alert(response.body.message);
            });

        },
        AddSubscriber: function(list, email, name){

            this.$http.post("subscriber/add", {list: list, email: email, name: name}).then(function(response) {
                var self = this;

                response.json().then(function(data){
                    vm.subscribers.push({name: data.name, emailaddress: data.emailaddress});

                    if (self.subscribers.length !== 0){
                        self.no_records_message = '';
                    }

                });

            }).catch(function(response){
                alert(response.body.message);
            });

        }
    }
});
