import type { RpcTransport } from '@protobuf-ts/runtime-rpc'

import { AdminServiceClient } from '../rpc/zhipin/api/admin/admin.client'
import {
    ListTodayPositionsReq,
    ListAllPositionsReq,
    GetPositionDetailReq,
    GetStatsReq,
} from '../rpc/zhipin/api/admin/admin'

// ========================== 响应类型定义 ==========================

export interface T岗位摘要 {
    s编号: string
    s岗位编号: string
    s岗位名称: string
    s公司名称: string
    s薪资范围: string
    n薪资下限: number
    n薪资上限: number
    s城市: string
    s招聘者: string
    b猎头: boolean
    n状态: number
    n匹配度: number
    s简历版本: string
    n最后沟通: string
    n最后方向: number
    s链接: string
}

export interface T匹配项 {
    s岗位要求: string
    n匹配状态: number
    s简历对应: string
    s补充说明: string
}

export interface T沟通项 {
    n事件类型: number
    s事件时间: string
    s消息内容: string
    n消息方向: number
}

export interface T岗位详情 {
    s岗位: T岗位摘要
    s匹配列表: T匹配项[]
    s沟通列表: T沟通项[]
    s岗位职责: string
    s岗位要求: string
    s备注: string
    s跳过原因: string
}

export interface T状态统计 {
    n状态: number
    n数量: number
}

// ========================== 转换函数 ==========================

function cnv岗位摘要(s: any): T岗位摘要 {
    return {
        s编号: s.id,
        s岗位编号: s.jobId,
        s岗位名称: s.title,
        s公司名称: s.company,
        s薪资范围: s.salaryRange,
        n薪资下限: s.salaryMin,
        n薪资上限: s.salaryMax,
        s城市: s.city,
        s招聘者: s.recruiter,
        b猎头: s.isHunter,
        n状态: s.status,
        n匹配度: s.matchRate,
        s简历版本: s.lastResume,
        n最后沟通: s.lastCommAt,
        n最后方向: s.lastCommDir,
        s链接: s.link,
    }
}

// ========================== SDK 类定义 ==========================

export class Sdk管理面板 {
    private rpc: AdminServiceClient

    constructor(transport: RpcTransport) {
        this.rpc = new AdminServiceClient(transport)
    }

    // 今日沟通的岗位
    async get今日岗位(): Promise<{ s岗位列表: T岗位摘要[]; n总数: number }> {
        const res = await this.rpc.listTodayPositions(ListTodayPositionsReq.create({}), {})
        return {
            s岗位列表: res.data.items.map(cnv岗位摘要),
            n总数: res.data.total,
        }
    }

    // 全部岗位列表（按状态筛选）
    async get全部岗位(status: number, page: number, pageSize: number): Promise<{ s岗位列表: T岗位摘要[]; n总数: number }> {
        const res = await this.rpc.listAllPositions(ListAllPositionsReq.create({ status, page, pageSize }), {})
        return {
            s岗位列表: res.data.items.map(cnv岗位摘要),
            n总数: res.data.total,
        }
    }

    // 岗位详情
    async get岗位详情(id: string): Promise<T岗位详情> {
        const res = await this.rpc.getPositionDetail(GetPositionDetailReq.create({ id }), {})
        const d = res.data
        return {
            s岗位: cnv岗位摘要(d.position),
            s匹配列表: d.matchItems.map(v => ({
                s岗位要求: v.requirement,
                n匹配状态: v.matchStatus,
                s简历对应: v.resumePoint,
                s补充说明: v.remark,
            })),
            s沟通列表: d.communications.map(v => ({
                n事件类型: v.eventType,
                s事件时间: v.eventTime?.seconds || '',
                s消息内容: v.content,
                n消息方向: v.direction,
            })),
            s岗位职责: d.duties,
            s岗位要求: d.requirements,
            s备注: d.notes,
            s跳过原因: d.skipReason,
        }
    }

    // 状态统计
    async get统计(): Promise<{ n总数: number; s状态统计: T状态统计[] }> {
        const res = await this.rpc.getStats(GetStatsReq.create({}), {})
        return {
            n总数: res.data.total,
            s状态统计: res.data.statusCounts.map(v => ({
                n状态: v.status,
                n数量: v.count,
            })),
        }
    }
}
