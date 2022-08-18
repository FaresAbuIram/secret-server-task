<script>
import axios from 'axios'
export default {
  data() {
    return {
      secret: '',
      expire: '60',
      numberOfViews: '5',
      required_feild: null,
      required_views: null
    }
  },
  methods: {
    handleSubmit() {
      this.required_feild = this.secret != '' ?
        '' : 'This is requied field'
      this.required_views = this.numberOfViews != '' ?
       '' : 'This is requied field'
        if(this.secret != '' && this.numberOfViews != ''){
          this.$isLoading(true)
         axios.post('http://localhost:9090/generate',{
                'data': this.secret,
                'expire': parseInt(this.expire),
                'views': parseInt(this.numberOfViews)
            }).then((response) => {
              this.$isLoading(false)
              this.$router.push({name: 'share', params: { token: response.data.data.replaceAll('.', '/')}}) 
            }).catch(error => {
              this.$isLoading(false)
              This.router.push({name: 'notfound'})
            })  
        }
      
    }
  }
}
</script>

<template>
    <div class="content">
        <div>
            <h1 class="title">Let's get your secret salted</h1>
        </div>
        <div class="menu">
            <div>
                <img src="https://www.saltify.io/assets/images/lock.png" alt="">
            </div>
            <div class="form">
               <form @submit.prevent="handleSubmit">
                    <label>Your Secret:</label>
                    <textarea v-model="secret" placeholder="add secret"></textarea>
                    <div v-if="required_feild" class="error">{{ required_feild }}</div>
                    
                    <label> Expires in:</label>
                    <select v-model="expire">
                            <optgroup label="keep for minutes">
                                <option value="5">5 mins</option>
                                <option value="15">10 mins</option>
                                <option value="30">15 mins</option>
                            </optgroup>
                            <optgroup label="keep for hours">
                                <option value="60">1 hour</option>
                                <option value="240">4 hours</option>
                                <option value="720">12 hours</option>
                            </optgroup>
                            <optgroup label="keep for days">
                                <option value="1440">1 day</option>
                                <option value="4320">3 days</option>
                                <option value="10080">7 days</option>
                            </optgroup>
                    </select>

                    <label>Number of Views</label>
                    <input type="number" v-model="numberOfViews" min="1" max="100">
                    <div v-if="required_views" class="error">{{ required_views }}</div>

                    <div class="submit">
                    <button>SHARE YOUR SECRET!</button>
                    </div>
                </form>
            </div>

        </div>
    </div>
</template>


<style scoped>
.content {
    display: flex;
    flex-direction: column;
    align-items: center;

}

.title {
    color: #545353;
    font-style: oblique;
    font-weight: 600;
    font-size: 50px;
}

.menu{
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
}

 form {
    min-width: 420px;
    margin: 30px auto;
    background: rgba(217, 210, 210, 0.412);
    text-align: left;
    padding: 40px;
    border-radius: 10px;
  }
  label {
    color: black;
    display: inline-block;
    margin: 25px 0 15px;
    font-size: 0.6em;
    text-transform: uppercase;
    letter-spacing: 1px;
    font-weight: bold;
  }
  textarea, select, input {
    display: block;
    padding: 10px 6px;
    width: 100%;
    box-sizing: border-box;
    border: none;
    border-bottom: 1px solid #ddd;
    color: #555;
    border-radius: 10px;
  }
  
  
  button {
    background: #0b6dff;
    border: 0;
    padding: 10px 20px;
    margin-top: 20px;
    color: white;
    border-radius: 20px;
  }
  .submit {
    text-align: center;
  }
  .error {
    color: #ff0062;
    margin-top: 10px;
    font-size: 0.8em;
    font-weight: bold;
  }

  img{
    width: 90%;
    min-width: 300px;
  }
</style>