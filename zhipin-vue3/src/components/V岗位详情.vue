<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { sdk管理面板, type T岗位详情 } from '../sdk'

const props = defineProps<{ positionId: string }>()

const detail = ref<T岗位详情 | null>(null)
const loading = ref(false)

const matchStatusMap: Record<number, { text: string; type: string }> = {
    0: { text: '未知', type: 'info' },
    1: { text: '✅ 匹配', type: 'success' },
    2: { text: '⚠️ 部分匹配', type: 'warning' },
    3: { text: '❌ 不匹配', type: 'danger' },
}

const eventTypeMap: Record<number, string> = {
    0: '未知', 1: '发消息', 2: '收消息', 3: '发简历',
    4: '开聊限制', 5: '安排面试', 6: '收到邀请', 7: '被拒绝',
}

async function loadDetail() {
    if (!props.positionId) return
    loading.value = true
    try {
        detail.value = await sdk管理面板.get岗位详情(props.positionId)
    } catch (e) {
        console.error(e)
    } finally {
        loading.value = false
    }
}

onMounted(loadDetail)
watch(() => props.positionId, loadDetail)
</script>

<template>
    <div v-loading="loading">
        <template v-if="detail">
            <!-- 基本信息 -->
            <el-descriptions :column="2" border style="margin-bottom: 20px">
                <el-descriptions-item label="公司">{{ detail.s岗位.s公司名称 }}</el-descriptions-item>
                <el-descriptions-item label="岗位">{{ detail.s岗位.s岗位名称 }}</el-descriptions-item>
                <el-descriptions-item label="薪资">{{ detail.s岗位.s薪资范围 }}</el-descriptions-item>
                <el-descriptions-item label="城市">{{ detail.s岗位.s城市 }}</el-descriptions-item>
                <el-descriptions-item label="招聘者">{{ detail.s岗位.s招聘者 }}</el-descriptions-item>
                <el-descriptions-item label="匹配度">{{ detail.s岗位.n匹配度 }}%</el-descriptions-item>
                <el-descriptions-item label="链接" :span="2">
                    <a :href="detail.s岗位.s链接" target="_blank" style="color: #409eff">{{ detail.s岗位.s链接 }}</a>
                </el-descriptions-item>
            </el-descriptions>

            <!-- 跳过原因 -->
            <el-alert v-if="detail.s跳过原因" :title="'跳过原因：' + detail.s跳过原因" type="warning" :closable="false" style="margin-bottom: 16px" />

            <!-- 匹配分析 -->
            <h3 style="margin: 16px 0 8px">匹配分析</h3>
            <el-table :data="detail.s匹配列表" stripe style="width: 100%; margin-bottom: 20px" v-if="detail.s匹配列表.length">
                <el-table-column label="岗位要求" min-width="200">
                    <template #default="{ row }">{{ row.s岗位要求 }}</template>
                </el-table-column>
                <el-table-column label="匹配" width="120">
                    <template #default="{ row }">
                        <el-tag :type="(matchStatusMap[row.n匹配状态] || matchStatusMap[0]).type as any">
                            {{ (matchStatusMap[row.n匹配状态] || matchStatusMap[0]).text }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="简历对应" min-width="200">
                    <template #default="{ row }">{{ row.s简历对应 }}</template>
                </el-table-column>
            </el-table>

            <!-- 沟通记录 -->
            <h3 style="margin: 16px 0 8px">沟通记录</h3>
            <el-timeline v-if="detail.s沟通列表.length">
                <el-timeline-item
                    v-for="(comm, idx) in detail.s沟通列表"
                    :key="idx"
                    :type="comm.n消息方向 === 0 ? 'primary' : 'success'"
                    :hollow="comm.n消息方向 === 1"
                >
                    <div style="display: flex; gap: 8px; align-items: baseline">
                        <el-tag size="small" :type="comm.n消息方向 === 0 ? 'primary' : 'success'">
                            {{ comm.n消息方向 === 0 ? '我→对方' : '对方→我' }}
                        </el-tag>
                        <span style="color: #909399; font-size: 12px">{{ eventTypeMap[comm.n事件类型] || '未知' }}</span>
                    </div>
                    <p style="margin: 4px 0 0; color: #303133">{{ comm.s消息内容 }}</p>
                </el-timeline-item>
            </el-timeline>
            <el-empty v-else description="暂无沟通记录" />

            <!-- 岗位职责和要求 -->
            <template v-if="detail.s岗位职责">
                <h3 style="margin: 16px 0 8px">岗位职责</h3>
                <pre style="white-space: pre-wrap; color: #606266; font-size: 13px; background: #f5f7fa; padding: 12px; border-radius: 4px">{{ detail.s岗位职责 }}</pre>
            </template>
            <template v-if="detail.s岗位要求">
                <h3 style="margin: 16px 0 8px">岗位要求</h3>
                <pre style="white-space: pre-wrap; color: #606266; font-size: 13px; background: #f5f7fa; padding: 12px; border-radius: 4px">{{ detail.s岗位要求 }}</pre>
            </template>
        </template>
    </div>
</template>
