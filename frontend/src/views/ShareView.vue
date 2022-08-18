<template>
    <div class="share">
        <h1>Yummy!</h1>
        <h4>The secret has successfully been encrypted and is now waiting for being shared:</h4>
        <textarea v-model="secret" disabled cols="30" rows="10"></textarea>
        <router-link to="/">Go back home</router-link>
    </div>
</template>

<script >
export default {
  data() {
    return {
        secret:  'http://' + location.host + '/check/' + this.$route.params.token
    }
  }
  }

</script>

<style scoped>
.share{
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    color: black;
    text-align: center;
    align-items: center;
    gap: 30px
}


textarea{
    display: block;
    padding: 10px 6px;
    box-sizing: border-box;
    min-width: 400px;
    width: 50%;
    border: none;
    border-bottom: 1px solid rgb(166, 160, 160);
    color: rgb(117, 111, 111);
    border-radius: 10px;
  }

 a {
    text-decoration: none;
    background: #e55039aa;
    color: #fff;
    padding: 12px 24px;
    display: inline-block;
    border-radius: 25px;
    font-size: 14px;
    text-transform: uppercase;
    transition: 0.4s;
    width: 200px;
}

a:hover {
    background: #e55039;
}

</style>