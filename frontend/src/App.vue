<script lang="ts" setup>
import { reactive } from 'vue'
import { GetPrograms, GetCheatSheet, GetVersion } from '../wailsjs/go/main/App'
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

const data = reactive({
  app_version: "",
  programs: [],
  currentContent: JSON.parse('{"Selecione um software no menu!": []}')
})

function changeContent(content: string) {
  data.currentContent = GetCheatSheet(content).then(result => {
    data.currentContent = JSON.parse(result)
  })
}

function AppVersion() {
  GetVersion().then(result => {
    data.app_version = result
  })
}
AppVersion()

function listPrograms() {
  GetPrograms().then(result => {
    data.programs = JSON.parse(result)
  })
}
listPrograms()
</script>

<template>
  <div id="app" class="app-container">
    <div class="vertical-menu">
      <a v-for="program in data.programs" :key="program" href="#" @click="changeContent(program)">{{ program }}</a>
      <br>
      <footer>
        <p class="app_footer">{{ data.app_version }}</p>
      </footer>
    </div>
    <div v-for="(value, key) in data.currentContent" class="content-area">
      <h4>{{ key }}</h4>
      <DataTable :value="value">
        <Column field="command"></Column>
        <Column field="about"></Column>
      </DataTable>
    </div>
  </div>
</template>

<style>
.app-container {
  height: 100vh;
}
</style>

<style scoped>
.vertical-menu {
  width: 200px;
  height: 100vh;
  padding: 10px;
  position: fixed;
  overflow-y: scroll;
  overflow-x: hidden;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
}

.vertical-menu a {
  display: block;
  padding: 10px;
  text-decoration: none;
}

.content-area {
  margin-left: 220px;
  padding-left: 20px;
  padding-top: 5px;
  width: calc(100% - 250px);
}
</style>