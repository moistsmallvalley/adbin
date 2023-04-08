export interface TableRequest {
  tableName: string;
}

export interface TableResponse {
  name: string;
  columns: Column[];
}

export interface RowsRequest {
  tableName: string;
}

export interface RowsResponse {
  columns: Column[];
  rows: Row[];
}

export interface Column {
  name: string;
  type: string;
  required: boolean;
  primaryKey: boolean;
  autoIncrement: boolean;
}

export interface Row {
  [key: string]: number | string | null;
}

export interface RowRequest {
  tableName: string;
  primaryKeys: string[];
}

export interface RowResponse {
  columns: Column[];
  row: Row;
}

export interface PostRowRequest {
  tableName: string;
  row: Row;
}

export interface PatchRowRequest {
  tableName: string;
  primaryKeys: string[];
  row: Row;
}
