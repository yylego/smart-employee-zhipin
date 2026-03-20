<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { sdk管理面板, type T岗位摘要 } from '../sdk'

const emit = defineEmits<{ (e: 'open-detail', id: string): void }>()

const positions = ref<T岗位摘要[]>([])
const total = ref(0)
const loading = ref(false)
const statusFilter = ref(-1)

const statusOptions = [
    { value: -1, label: '全部' },
    { value: 1, label: '待处理' },
    { value: 2, label: '已跳过' },
    { value: 3, label: '开聊限制' },
    { value: 4, label: '已发消息' },
    { value: 5, label: '已回复' },
    { value: 6, label: '已发简历' },
    { value: 7, label: '面试中' },
    { value: 8, label: '已拿到' },
    { value: 9, label: '已拒绝' },
    { value: 10, label: '不再联系' },
]

const statusMap: Record<number, string> = {}
statusOptions.forEach(o => { statusMap[o.value] = o.label })

async function loadData() {
    loading.value = true
    try {
        const res = await sdk管理面板.get全部岗位(statusFilter.value, 1, 500)
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
        <div style="margin-bottom: 12px; display: flex; align-items: center; gap: 12px">
            <span style="color: #606266">状态筛选：</span>
            <el-select v-model="statusFilter" @change="loadData" style="width: 140px">
                <el-option v-for="opt in statusOptions" :key="opt.value" :value="opt.value" :label="opt.label" />
            </el-select>
            <span style="color: #909399; margin-left: auto">共 {{ total }} 个</span>
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
            <el-table-column prop="s城市" label="城市" width="80" />
            <el-table-column label="猎头" width="60">
                <template #default="{ row }">
                    {{ row.b猎头 ? '是' : '否' }}
                </template>
            </el-table-column>
            <el-table-column label="操作" width="80">
                <template #default="{ row }">
                    <el-button type="primary" link @click="emit('open-detail', row.s编号)">详情</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
