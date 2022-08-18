<template>
    <div class="ok">
        <div class="content" v-if="message === 'success'">
            <div class="success" >
                <h1>SOMEONE GAVE YOU THAT URL?</h1>
                <img src="/salt.png" alt="">
                <button @click="showData">REVEAL THE SECRET NOW!</button>
            </div>
            <div class="data" v-if="clicked">
                <label>Your Secret:</label>
                <textarea v-model="data" disabled cols="30" rows="10"></textarea>

                <label>Expiration Date:</label>

                <input v-model="date"   type="datetime-local" disabled>
               
                <label>Remaning Views:</label>
                <input v-model="views" type="text" disabled>
            </div>

        </div>
        <div class="wrong" v-if="message === 'failed'">
                <h1>SOMETHING WENT WRONG!</h1>
                <div>
                    <h4>Looks like it has never existed a secret {{token}} </h4>
                    <h5>Maybe it has already been viewed too many times or has expired in the meantime.</h5>
                </div>
                <img src="/salt.png" alt="">
                <router-link to="/">Go back home</router-link>
        </div>

    </div>
</template>

<script >
import axios from 'axios'
export default {
    data() {
        return {
            token: this.$route.params.token + '.' + this.$route.params.token1 + '.' + this.$route.params.token2,
            data: null,
            message: 'loading',
            clicked: null,
            date: null,
            views: null
        }
    },
    created() {
        this.$isLoading(true)
        axios.post(`http://localhost:9090/get/${this.token}`)
        .then(response => {
            this.data = response.data.data
            if(response.data && response.data.message === 'success'){
                this.message = 'success'
                this.date = response.data.object.ExpireDate.replace('T', '')
                this.date = response.data.object.ExpireDate.replace('Z', '')
                this.views = response.data.object.Views - 1
                
            }
            else{
                this.message = 'failed'
            }
            this.$isLoading(false)
        }).catch(error => {
            this.message = 'failed'
            this.$isLoading(false)
        })
    },
    methods: {
        showData(){
            this.clicked = true
        }
    },
}
</script>

<style scoped>
.content{
    display: flex;
    color: black;
    justify-content: space-around;
}
.success{
    display: flex;
    flex-direction: column;
    align-items:center;
    color: black;
    font-weight: bolder;
    gap: 30px
}

button, a {
    text-align: center;
    text-decoration: none;
    background: #e55039aa;
    color: #fff;
    padding: 12px 24px;
    display: inline-block;
    border-radius: 25px;
    font-size: 14px;
    text-transform: uppercase;
    transition: 0.4s;
    width: 300px;
    border: none;
}
button:hover, a:hover {
    background: #e55039;
}
 
 img{
    width: 200px;
    height: 250px;
 }

.data{
    display: flex;
    flex-direction: column;
    text-align: center;
    justify-content: center;
}
label {
    color: black;
    display: inline-block;
    margin: 25px 0 15px;
    font-size: 15px;
    text-transform: uppercase;
    letter-spacing: 1px;
    font-weight: bold;
  }
textarea{
    display: block;
    padding: 10px 6px;
    box-sizing: border-box;
    min-width: 180px;
    max-width: 300px;
    max-height: 300px;
    width: 50%;
    border: none;
    border-bottom: 1px solid rgb(166, 160, 160);
    color: rgb(117, 111, 111);
    border-radius: 10px;
  }
  input{
     display: block;
    padding: 10px 6px;
    box-sizing: border-box;
    width: 180px;
    border: none;
    border-bottom: 1px solid rgb(166, 160, 160);
    color: rgb(117, 111, 111);
    border-radius: 10px;
  }
  .wrong{
    display: flex;
    flex-direction: column;
    font-weight: bolder;
    gap: 30px;
    align-items:center;
    color: black;

  }
  h4{
    overflow: hidden;
    display: inline-block;
    text-overflow: ellipsis;
    white-space: nowrap;
  	width:340px;
  }
</style>