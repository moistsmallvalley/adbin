import { Column, Row } from "./payloads";

export function parseColumnValue(col: Column, s: string): string | number {
  if (isIntColumn(col)) {
    return parseInt(s, 10);
  }
  if (isFloatColumn(col)) {
    return parseFloat(s);
  }
  return s;
}

function isIntColumn(col: Column): boolean {
  return [
    "integer",
    "int",
    "smallint",
    "tinyint",
    "mediumint",
    "bigint",
    "bit",
  ].includes(col.type);
}

function isFloatColumn(col: Column): boolean {
  return ["decimal", "numeric", "float", "double"].includes(col.type);
}
