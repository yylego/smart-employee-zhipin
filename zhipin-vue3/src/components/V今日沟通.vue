<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { sdk管理面板, type T岗位摘要 } from '../sdk'

const emit = defineEmits<{ (e: 'open-detail', id: string): void }>()

const positions = ref<T岗位摘要[]>([])
const total = ref(0)
const loading = ref(false)

const statusMap: Record<number, string> = {
    0: '未知', 1: '待处理', 2: '已跳过', 3: '开聊限制',
    4: '已发消息', 5: '已回复', 6: '已发简历',
    7: '面试中', 8: '已拿到', 9: '已拒绝', 10: '不再联系',
}

async function loadData() {
    loading.value = true
    try {
        const res = await sdk管理面板.get今日岗位()
        positions.value = res.s岗位列表
        total.value = res.n总数
    } catch (e) {
        console.error(e)
    } finally {
        loading.value = false
    }
}

onMounted(loadData)
</script>

<template>
    <div>
        <div style="margin-bottom: 12px; color: #606266">
            今日沟通岗位：{{ total }} 个
        </div>

        <el-table :data="positions" v-loading="loading" stripe style="width: 100%">
            <el-table-column prop="s公司名称" label="公司" width="140" />
            <el-table-column prop="s岗位名称" label="岗位" width="160" />
            <el-table-column prop="s薪资范围" label="薪资" width="130" />
            <el-table-column label="状态" width="100">
                <template #default="{ row }">
                    {{ statusMap[row.n状态] || '未知' }}
                </template>
            </el-table-column>
            <el-table-column label="匹配度" width="80">
                <template #default="{ row }">
                    {{ row.n匹配度 }}%
                </template>
            </el-table-column>
            <el-table-column prop="s招聘者" label="招聘者" width="120" />
            <el-table-column label="操作" width="80">
                <template #default="{ row }">
                    <el-button type="primary" link @click="emit('open-detail', row.s编号)">详情</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
