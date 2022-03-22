<template>
    <b-card title="Deteksi Kualitas Tanaman Padi 🌾">
        <b-card-text>Input dalam Berupa Gambar / Image</b-card-text>
        <b-form @submit.prevent="onUpload">
            <b-form-file
            v-model="image"
            :state="Boolean(image)"
            placeholder="Choose a file or drop it here..."
            drop-placeholder="Drop file here..." @change="onChange"
            />
           
            <b-row style="margin-top: 20px;">
                <b-col cols-md="4" cols="12">
                <div v-if="isLoaded">
                    <b-button
                        v-ripple.400="'rgba(255, 255, 255, 0.15)'"
                        type="submit"
                        variant="primary"
                        class="mr-1"
                    >
                        Submit
                    </b-button>
                </div>
                <div v-else>
                    Loading....
                </div>
            </b-col>
            </b-row>
        </b-form>
    </b-card>
</template>

<script>
import axios from "axios"
import Ripple from 'vue-ripple-directive'
import { BCard, BCardText, BFormFile, BButton, BForm, BCol, BRow } from 'bootstrap-vue';

export default {
  components: {
    BCard,
    BCardText,
    BFormFile,
    BButton,
    BForm,
    BCol,
    BRow
  },
  data(){
      return {
          image: null,
          isLoaded: true,
      };
  },
  methods: {
      onChange(event){
          this.image = event.target.files[0]
      },
      onUpload(){
          const formData = new FormData()
          formData.append('image', this.image, this.image.name)
          this.isLoaded = false;
          axios.post('http://localhost:9000/image-services', formData, {
          }).then((response) => {
              console.log(response)
              if (response.status === 200){
                  this.$router.push({name: 'result', params: {urlofimg: response.data.data.imageUrl}})
              }
          })
      }
  },
  directives: {
      Ripple
  }
}

</script>

<style>

</style>
