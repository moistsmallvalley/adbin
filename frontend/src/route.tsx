import React from "react";
import { createBrowserRouter, generatePath } from "react-router-dom";
import App from "./App";
import { NewRowForm } from "./NewRowForm";
import { RowForm } from "./RowForm";
import { TableView } from "./TableView";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "tables/:tableName/rows",
        element: <TableView />,
      },
      {
        path: "tables/:tableName/new",
        element: <NewRowForm />,
      },
      {
        path: "tables/:tableName/rows/:keys",
        element: <RowForm />,
      },
    ],
  },
]);

export function rootPath() {
  return "/";
}

export function rowsPath(tableName: string) {
  return generatePath("/tables/:tableName/rows", { tableName });
}

export function newRowPath(tableName: string) {
  return generatePath("/tables/:tableName/new", { tableName });
}

export function rowPath(tableName: string, keys: string) {
  return generatePath("/tables/:tableName/rows/:keys", { tableName, keys });
}
