<script lang="ts" setup>
import { reactive } from 'vue'
import { GetPrograms, GetCheatSheet } from '../wailsjs/go/main/App'

const data = reactive({
  programs: [],
  currentContent: JSON.parse('{"Selecione um software no menu!": []}')
})

function changeContent(content: string) {
  data.currentContent = GetCheatSheet(content).then(result => {
    data.currentContent = JSON.parse(result)
  })
}

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
    </div>
    <div v-for="(value, key) in data.currentContent" class="content-area">
      <h4 id="tabela1">{{ key }}</h4>
      <table v-for="vvalue in value">
        <tr v-for="(chv, vle) in vvalue">
          <td>{{ chv }}</td>
          <td>{{ vle }}</td>
        </tr>
      </table>
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
  background-color: #f8f9fa;
  padding: 10px;
  position: fixed;
  overflow-y: scroll;
  overflow-x: hidden;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
}

.vertical-menu a {
  display: block;
  padding: 10px;
  color: #333;
  text-decoration: none;
}

.vertical-menu a:hover {
  background-color: #ddd;
}

.content-area {
  margin-left: 220px;
  padding: 20px;
  width: calc(100% - 220px);
}
</style>