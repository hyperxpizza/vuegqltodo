<template>
    <div>
        <!-- Form -->
        <form id="addForm">
            <label for="First Name">First Name</label>
            <input type="text" v-model="firstName" required/>
            <br>
            <label for="Last Name">Last Name</label>
            <input type="text" v-model="lastName" required/>
            <br>
            <label for="Email">Email</label>
            <input type="email" v-model="email" required/>
            <br>
            <label for="Phone">Phone</label>
            <input type="tel" v-model="phone" required/>
        </form>
        <!--Buttons-->
        
        <button type="submit" form="addForm" class="actionButton" id="addButton" @click="addContact">Add</button>
        <button class="actionButton" id="clearButton">Clear</button>

    </div>
</template>

<script>
import gql from 'graphql-tag';

const ADD_CONTACT =  gql`
    mutation($firstName: String!, $lastName: String!, $email: String!, $phone: String!){
        createContact(input:{
            firstName: $firstName,
            lastName: $lastName,
            email: $email,
            phone: $phone
        }){
            firstName
            lastName
            email
            phone
            createdAt
            updatedAt
        }
    }
`;

export default {
    name:'AddContactForm',
    data() {
        return {
            firstName: '',
            lastName: '',
            email: '',
            phone: ''
        }
    },
    methods: {
        async addContact(){
            const firstName = this.firstName;
            const lastName = this.lastName;
            const email = this.email;
            const phone = this.phone;

            await this.$apollo.mutate({
                mutation: ADD_CONTACT,
                variables: {
                    firstName,
                    lastName,
                    email,
                    phone
                }
            }).then(data => {
                console.log(data);
            }).catch(error => {
                console.log(error);
            });
        }
    }
}
</script>

<style scoped>
    form{
        width: 80%;
        margin: 0 auto;
    }

    label, input {
        display: inline-block;
    }

    input{
        padding: 6px 20px;
        margin: 8px 0;
        box-sizing: border-box;
    }

    #addButton{
        background-color: #4CAF50;
        margin-right: 5px;
    }

    #clearButton{
        background-color: #008CBA;
        margin-left: 5px;
    }
</style>