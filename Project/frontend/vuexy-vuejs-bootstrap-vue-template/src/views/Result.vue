<template>
  <b-card title="Results">
    <b-card-text>Chocolate sesame snaps pie carrot cake pastry pie lollipop muffin. Carrot cake dragée chupa chups jujubes. Macaroon liquorice cookie wafer tart marzipan bonbon. Gingerbread jelly-o dragée chocolate.</b-card-text>
    <div class="container">
        <div class="row text-center">
            <div class="col">
            </div>
            <div class="col-8">
                <img v-bind:src=this.$route.params.urlofimg fluid-grow alt="Fluid-grow image" width="500px">
                <br>
                <a class="btn btn-primary" style="margin-top:20px;" href="/" role="button">Back</a>
            </div>
            <div class="col">
            </div>
        </div>
    </div>
  </b-card>
</template>

<script>
import { BCard, BCardText, BRow, BContainer} from 'bootstrap-vue'
export default {
  components: {
    BCard,
    BCardText,BRow, BContainer
  },
}

</script>

<style>

</style>
