import Vue from 'vue'
import App from './App.vue'
import { ApolloClient } from 'apollo-client'
import VueApollo from "vue-apollo"
import { createHttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'

Vue.config.productionTip = false

const httpLink = createHttpLink({
  uri: "http://localhost:8888/query",
});

const cache = new InMemoryCache()

const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
})

Vue.use(VueApollo)

const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
})

new Vue({
  render: h => h(App),
  apolloProvider,
}).$mount('#app')
