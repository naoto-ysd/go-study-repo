<template>
  <main class="container">
    <h1>地域別テーブル</h1>
    <div class="table-vert-scroll" role="region" aria-label="地域別テーブル（縦スクロール）">
      <table>
        <caption>{{ columns.join('/') }} の値</caption>
        <thead>
          <tr>
            <th v-for="col in columns" :key="col" scope="col"
                @click="openModal(col)"
                @keydown.enter.prevent="openModal(col)"
                @keydown.space.prevent="openModal(col)"
                tabindex="0"
                class="clickable-header"
                :aria-label="col + ' をクリックで詳細モーダルを表示'">
              {{ col }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(r,i) in rows" :key="i">
            <td v-for="(val,j) in r" :key="j">{{ val }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
  <AppDialog :label="selectedHeader + ' ヘッダー情報'" v-model:show="isDialogOpen" @close="closeModal">
    <p style="margin-top:0; font-weight:600">{{ selectedHeader }} のヘッダーがクリックされました</p>
    <p style="font-size:.9rem; color:#555">ここに {{ selectedHeader }} に関する追加情報を表示できます。</p>
    <div style="text-align:right; margin-top:16px;">
      <button type="button" @click="closeModal" class="btn-close">閉じる</button>
    </div>
  </AppDialog>
</template>

<script setup>
import { ref } from 'vue';
import AppDialog from './components/AppDialog.vue';

const columns = ref(["岐阜","東京","名古屋","大阪","福岡"]);
const rows = ref([
  [100,200,300,400,500],
  [600,700,800,900,1000],
  [600,700,800,900,1000],
  [600,700,800,900,1000],
  [600,700,800,900,1000],
  [600,700,800,900,1000],
  [600,700,800,900,1000]
]);
const selectedHeader = ref('');
const isDialogOpen = ref(false);
const openModal = (col) => { selectedHeader.value = col; isDialogOpen.value = true; };
const closeModal = () => { isDialogOpen.value = false; };
</script>

<style>
</style>
