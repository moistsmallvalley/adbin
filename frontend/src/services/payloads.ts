export interface TableResponse {
  columns: string[];
  rows: TableRow[];
}

export interface TableRow {
  [key: string]: any;
}
