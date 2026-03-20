<script setup lang="ts">
import { ref } from 'vue'
import V今日沟通 from './components/V今日沟通.vue'
import V全部岗位 from './components/V全部岗位.vue'
import V岗位详情 from './components/V岗位详情.vue'

const activeTab = ref('today')
const detailId = ref('')

function openDetail(id: string) {
    detailId.value = id
    activeTab.value = 'detail'
}

function backToList() {
    detailId.value = ''
    activeTab.value = 'today'
}
</script>

<template>
    <div class="app">
        <div class="app-header">
            <h1>求职进度管理</h1>
        </div>

        <el-card class="main-card" shadow="always">
            <el-tabs v-model="activeTab" v-if="activeTab !== 'detail'">
                <el-tab-pane label="今日沟通" name="today">
                    <V今日沟通 @open-detail="openDetail" />
                </el-tab-pane>
                <el-tab-pane label="全部岗位" name="all">
                    <V全部岗位 @open-detail="openDetail" />
                </el-tab-pane>
            </el-tabs>

            <div v-if="activeTab === 'detail'">
                <el-button @click="backToList" style="margin-bottom: 16px">← 返回列表</el-button>
                <V岗位详情 :position-id="detailId" />
            </div>
        </el-card>
    </div>
</template>

<style scoped>
.app {
    max-width: 1100px;
    margin: 0 auto;
    padding: 24px 16px;
}
.app-header {
    text-align: center;
    margin-bottom: 16px;
}
h1 {
    margin: 0;
    color: #303133;
}
.main-card {
    border-radius: 12px;
}
</style>
