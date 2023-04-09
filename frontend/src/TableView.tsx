import React from "react";
import {
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Grid,
  Button,
  Stack,
} from "@mui/material";
import { skipToken } from "@reduxjs/toolkit/dist/query";
import { NavLink, useParams } from "react-router-dom";
import { useGetRowsQuery } from "./services/tables";
import EditIcon from "@mui/icons-material/Edit";
import { Column, Row } from "./services/payloads";
import { newRowPath, rowPath } from "./route";

export function TableView() {
  const { tableName } = useParams();
  const { data, error, isLoading } = useGetRowsQuery(
    tableName ? { tableName } : skipToken
  );

  return error ? (
    <>Network Error</>
  ) : isLoading ? (
    <>Loading...</>
  ) : data && tableName ? (
    <Stack>
      <Grid container sx={{ my: 2 }}>
        <NavLink to={newRowPath(tableName)}>
          <Button variant="outlined">New Item</Button>
        </NavLink>
      </Grid>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell />
              {data.columns.map((column) => (
                <TableCell key={column.name}>{column.name}</TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data.rows.map((row, i) => (
              <TableRow key={i}>
                <TableCell>
                  <NavLink
                    to={rowPath(tableName, primaryKeyPath(row, data.columns))}
                  >
                    <EditIcon />
                  </NavLink>
                </TableCell>
                {data.columns.map((column) => (
                  <TableCell key={column.name}>{row[column.name]}</TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Stack>
  ) : null;
}

function primaryKeyPath(row: Row, columns: Column[]) {
  return columns
    .filter((col) => col.primaryKey)
    .map((col) => encodeURIComponent(row[col.name]?.toString() ?? ""))
    .join("/");
}
