<template>
  <b-card-code title="List Hasil Deteksi">

    <!-- search input -->
    <div class="custom-search d-flex justify-content-end">
      <b-form-group>
        <div class="d-flex align-items-center">
          <label class="mr-1">Search</label>
          <b-form-input
            v-model="searchTerm"
            placeholder="Search"
            type="text"
            class="d-inline-block"
          />
        </div>
      </b-form-group>
    </div>

    <!-- table -->
    <vue-good-table
      :columns="columns"
      :rows="rows"
      :search-options="{
        enabled: true,
        externalQuery: searchTerm }"
      :sort-options="{
            enabled: true,
            initialSortBy: {field: 'timestamp', type: 'asc'}
        }"
      :pagination-options="{
        enabled: true,
        perPage:pageLength
      }"
      styleClass="vgt-table bordered container"
    >
    
      <template
        slot="table-row"
        slot-scope="props"
      >
        <!-- Column: Name -->
        <span
          v-if="props.column.field === 'imgUrlData'"
          class="text-nowrap"
        >
            <img v-bind:src=props.row.imgUrlData fluid-grow alt="Fluid-grow image" width="200px" @click="showModal(props.row.imgUrlData)">
            <div id="app">
                <b-modal id="showImageModal" v-b-modal.modal-xl hide-footer no-close-on-backdrop>
                    <template #modal-title>
                        Detail of Images
                    </template>
                    <div class="d-block text-center">
                        <img v-bind:src="modal_image_path" width="400px"/>
                    </div>
                    <b-button style="margin-top: 20px;" variant="danger" block @click="$bvModal.hide('showImageModal')">Close Me</b-button>
                </b-modal>
            </div>
        </span>

        <!-- Column: Action -->
        <span v-else-if="props.column.field === 'action'">
          <span>
            <b-dropdown
              variant="link"
              toggle-class="text-decoration-none"
              no-caret
            >
              <template v-slot:button-content>
                <feather-icon
                  icon="MoreVerticalIcon"
                  size="16"
                  class="text-body align-middle mr-25"
                />
              </template>
              <b-dropdown-item>
                <feather-icon
                  icon="Edit2Icon"
                  class="mr-50"
                />
                <span>Edit</span>
              </b-dropdown-item>
              <b-dropdown-item>
                <feather-icon
                  icon="TrashIcon"
                  class="mr-50"
                />
                <span>Delete</span>
              </b-dropdown-item>
            </b-dropdown>
          </span>
        </span>

        <!-- Column: Common -->
        <span v-else>
          {{ props.formattedRow[props.column.field] }}
        </span>
      </template>

      <!-- pagination -->
      <template
        slot="pagination-bottom"
        slot-scope="props"
      >
        <div class="d-flex justify-content-between flex-wrap">
          <div class="d-flex align-items-center mb-0 mt-1">
            <span class="text-nowrap ">
              Showing 1 to
            </span>
            <b-form-select
              v-model="pageLength"
              :options="['2','4','6', '8', '10']"
              class="mx-1"
              @input="(value)=>props.perPageChanged({currentPerPage:value})"
            />
            <span class="text-nowrap"> of {{ props.total }} entries </span>
          </div>
          <div>
            <b-pagination
              :value="1"
              :total-rows="props.total"
              :per-page="pageLength"
              first-number
              last-number
              align="right"
              prev-class="prev-item"
              next-class="next-item"
              class="mt-1 mb-0"
              @input="(value)=>props.pageChanged({currentPage:value})"
            >
              <template #prev-text>
                <feather-icon
                  icon="ChevronLeftIcon"
                  size="18"
                />
              </template>
              <template #next-text>
                <feather-icon
                  icon="ChevronRightIcon"
                  size="18"
                />
              </template>
            </b-pagination>
          </div>
        </div>
      </template>
    </vue-good-table>
  </b-card-code>
</template>

<script>
import BCardCode from '@core/components/b-card-code/BCardCode.vue'
import {
  BAvatar, BBadge, BPagination, BFormGroup, BFormInput, BFormSelect, BDropdown, BDropdownItem, BModal, BButton
} from 'bootstrap-vue'
import { VueGoodTable } from 'vue-good-table'
import { codeBasic } from './code'

export default {
  components: {
    BCardCode,
    VueGoodTable,
    BAvatar,
    BBadge,
    BPagination,
    BFormGroup,
    BFormInput,
    BFormSelect,
    BDropdown,
    BDropdownItem,
    BModal,
    BButton,
    codeBasic
  },
  data() {
    return {
      modal_image_path: '',
      pageLength: 3,
      dir: false,
      codeBasic,
      columns: [
        {
          label: 'TimeStamp',
          field: 'timestamp',
        },
        {
          label: 'Image',
          field: 'imgUrlData',
          sortable: false,
        },
        {
          label: 'Action',
          field: 'action',
        },
      ],
      rows: [],
      searchTerm: '',
    }
  },
  methods: {
    showModal(imgPath) {
      this.modal_image_path = imgPath
      this.$bvModal.show('showImageModal')
    }
  },
  computed: {
     direction() {
      // eslint-disable-next-line vue/no-side-effects-in-computed-properties
      this.dir = true
      return this.dir
    },

  },
  created() {
    this.$http.get('http://localhost:9000/imageDetectionData')
      .then(res => { this.rows = res.data.data })
  },
}
</script>
<style>
.close {display: none;}
</style>