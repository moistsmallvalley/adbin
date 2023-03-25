import {
  Paper,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Box,
} from "@mui/material";
import { skipToken } from "@reduxjs/toolkit/dist/query";
import { useParams } from "react-router-dom";
import { useGetTableQuery } from "./services/tables";

export interface TableProps {
  name: string;
}

export function TableView() {
  const { table } = useParams();
  const { data, error, isLoading } = useGetTableQuery(table ?? skipToken);

  return error ? (
    <>Network Error</>
  ) : isLoading ? (
    <>Loading...</>
  ) : data ? (
    <TableContainer>
      <Table>
        <TableHead>
          <TableRow>
            {data.columns.map((column) => (
              <TableCell key={column}>{column}</TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {data.rows.map((row) => (
            <TableRow key={row.id}>
              {data.columns.map((column) => (
                <TableCell key={column}>{row[column]}</TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  ) : null;
}
