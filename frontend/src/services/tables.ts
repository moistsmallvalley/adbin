import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {
  PatchRowRequest,
  PostRowRequest,
  RowRequest,
  RowResponse,
  RowsRequest,
  RowsResponse,
  TableRequest,
  TableResponse,
} from "./payloads";

const apiHost = import.meta.env.API_HOST ?? "http://localhost:8080";

export const tablesApi = createApi({
  reducerPath: "tablesApi",
  baseQuery: fetchBaseQuery({ baseUrl: apiHost + "/api/" }),
  endpoints: (builder) => ({
    getTables: builder.query<string[], void>({
      query: () => "tables",
    }),
    getTable: builder.query<TableResponse, TableRequest>({
      query: (req) => `tables/${req.tableName}`,
    }),
    getRows: builder.query<RowsResponse, RowsRequest>({
      query: (req) => `tables/${req.tableName}/rows`,
    }),
    getRow: builder.query<RowResponse, RowRequest>({
      query: (req) =>
        `tables/${req.tableName}/rows/${req.primaryKeys
          .map(encodeURIComponent)
          .join("/")}`,
    }),
    postRow: builder.mutation<RowResponse, PostRowRequest>({
      query: (req) => ({
        url: `tables/${req.tableName}/rows`,
        method: "POST",
        body: req.row,
      }),
    }),
    patchRow: builder.mutation<RowResponse, PatchRowRequest>({
      query: (req) => ({
        url: `tables/${req.tableName}/rows/${req.primaryKeys
          .map(encodeURIComponent)
          .join("/")}`,
        method: "PATCH",
        body: req.row,
      }),
    }),
  }),
});

export const {
  useGetTablesQuery,
  useGetTableQuery,
  useGetRowsQuery,
  useGetRowQuery,
  usePostRowMutation,
  usePatchRowMutation,
} = tablesApi;
