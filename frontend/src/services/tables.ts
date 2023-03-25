import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { TableResponse } from "./payloads";

const apiHost = import.meta.env.API_HOST ?? "http://localhost:8080";

export const tablesApi = createApi({
  reducerPath: "tablesApi",
  baseQuery: fetchBaseQuery({ baseUrl: apiHost + "/api/" }),
  endpoints: (builder) => ({
    getTables: builder.query<string[], void>({
      query: () => "tables",
    }),
    getTable: builder.query<TableResponse, string>({
      query: (name) => `tables/${name}`,
    }),
  }),
});

export const { useGetTablesQuery, useGetTableQuery } = tablesApi;
