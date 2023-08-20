import type {
  ListReponse,
  PageReply,
  PageReq,
  Query,
  Response,
} from "@/apis/type";
import type { GroupItem, GroupItemRequest, SimpleItem } from "@/apis/prom/prom";

// 规则组列表请求参数
export type ListGroupRequest = {
  query: Query;
  group?: GroupItemRequest;
};

export type ListSimpleGroupRequest = {
  page: PageReq;
  keyword?: string;
};

export const defaultListGroupRequest: ListGroupRequest = {
  query: {
    page: {
      current: 1,
      size: 10,
    },
  },
};

// 规则组列表响应参数
export type ListGroupReply = ListReponse & {
  groups: GroupItem[];
};

// 规则组列表响应参数
export type ListSimpleGroupReply = Response & {
  page: PageReply;
  groups: SimpleItem[];
};

// 规则组详情响应参数
export type GroupDetailReply = Response & {
  group: GroupItem;
};

// 规则组创建参数
export type GroupCreateItem = {
  name: string;
  remark: string;
  categoriesIds: number[];
};

// 规则组更新参数
export type GroupUpdateItem = {
  name: string;
  remark: string;
  categoriesIds: number[];
};
