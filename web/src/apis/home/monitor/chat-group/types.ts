import {
    AlarmApp,
    AlarmStatus,
    NotifyApp,
    NotifyTemplateType,
    PageReqType,
    PageResType,
    Status
} from '@/apis/types'
import { UserSelectItem } from '../../system/user/types'

interface ChatGroupItem {
    id: number
    name: string
    remark: string
    createdAt: number
    updatedAt: number
    hook: string
    status: number
    app: number
    hookName: string
    title: string
    template: string
    secret: string
    createUser?: UserSelectItem
}

interface GetChatGroupDetailRequest {
    id: number
}

interface GetChatGroupDetailResponse {
    detail: ChatGroupItem
}

interface DeleteChatGroupRequest {
    id: number
}

interface DeleteChatGroupResponse {
    id: number
}

interface UpdateChatGroupRequest {
    id: number
    name: string
    remark: string
    hookName: string
    title: string
    template: string
}

interface UpdateChatGroupResponse {
    id: number
}

interface ListChatGroupRequest {
    page: PageReqType
    keyword?: string
    status?: AlarmStatus
}

interface ListChatGroupResponse {
    list?: ChatGroupItem[]
    page: PageResType
}

interface ChatGroupSelectItem {
    value: number
    app: NotifyApp
    label: string
    status: number
}

interface SelectChatGroupRequest {
    page: PageReqType
    keyword?: string
    status?: Status
}

interface SelectChatGroupResponse {
    list: ChatGroupSelectItem[]
    page: PageResType
}

interface CreateChatGroupRequest {
    name: string
    remark: string
    hook: string
    app: AlarmApp
    hookName: string
    secret?: string
    template: string
}

interface CreateChatGroupResponse {
    id: number
}

interface TestTemplateRequest {
    notifyType: NotifyTemplateType
    template: string
    strategyId: number
}

interface TestTemplateResponse {
    msg: string
}

export const defaultSelectChatGroupReques: SelectChatGroupRequest = {
    page: {
        curr: 1,
        size: 10
    }
}

export const defaultListChatGroupRequest: ListChatGroupRequest = {
    page: {
        curr: 1,
        size: 10
    }
}

export type {
    ListChatGroupRequest,
    ListChatGroupResponse,
    SelectChatGroupRequest,
    SelectChatGroupResponse,
    UpdateChatGroupRequest,
    UpdateChatGroupResponse,
    DeleteChatGroupRequest,
    DeleteChatGroupResponse,
    GetChatGroupDetailRequest,
    GetChatGroupDetailResponse,
    CreateChatGroupRequest,
    CreateChatGroupResponse,
    ChatGroupItem,
    ChatGroupSelectItem,
    TestTemplateRequest,
    TestTemplateResponse
}
