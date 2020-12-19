<template>
  <div id="app">
    <table border="1" width="100%" style="border-collapse: collapse;">
      <tr>
        <th>First Name</th>
        <th>Last Name</th>
        <th>Email</th>
        <th>Phone</th>
        <th>Actions</th>
      </tr>

      <tr v-for ></tr>
    </table>
  </div>
</template>

<script>
import gql from 'graphql-tag';

export default {
  name: 'App',
  data(){
    return {
      id: null,
      firstName: '',
      lastName: '',
      email: '',
      phone: ''
    }
  },
  apollo: {
    contacts: gql`query {
      getAllContacts {
        id,
        firstName,
        lastName,
        email,
        phone
      }
    }`
  },
  methods: {
    createContact(firstName, lastName, email, phone){

      const currentTimestamp = new Date("2016-07-27T07:45:00Z");
      console.log(currentTimestamp);

      console.log(`Create contact: ${email}`)
      this.$apollo.mutate({
        mutation: gql`mutation createContact($firstName: String!, $lastName: String!, $email: String!, $phone: String!){
          createContact(input: {
            firstName: $firstName,
            lastName: $lastName,
            email: $email,
            phone: $phone,
          })
        }`,
        variables: {
          firstName: firstName,
          lastName: lastName,
          email: email,
          phone: phone,
        }
      })
      location.reload();
    },
    updateContact(id, firstName, lastName, email, phone){

      const currentTimestamp = new Date("2016-07-27T07:45:00Z");

      console.log(`Update contact: #${id}`);
      this.$apollo.mutate({
        mutation: gql`mutation updateContact($id: ID!, $firstName: String!, $lastName: String!, $email: String!, $phone: String!){
          updateContact(input: {
            id: $id,
            firstName: $firstName,
            lastName: $lastName,
            email: $email,
            phone: $phone
          })
        }
        `,
        variables: {
          id: id,
          firstName: firstName,
          lastName: lastName,
          email: email,
          phone: phone,
        }
      })

      location.reload();

    },
    deleteContact(id){
      console.log(`Delete contact: # ${id}`)
      
      this.$apollo.mutate({
        mutation: gql`mutation deleteContact($id: ID!){
          deleteContact(id: $id)
          }`,
          variables:{
            id: id,
          }
        })
      location.reload();
    },
    selectContact(contact){
      this.id = contact.id; 
      this.firstName = contact.firstName;
      this.lastName = contact.lastName;
      this.email = contact.emaill;
      this.phone = contact.phone;
    },
    cleanForm(){
      this.id = null;
      this.firstName = '';
      this.lastName = '';
      this.email = '';
      this.phone = '';
    }
  }  
}
</script>

<style>

</style>
