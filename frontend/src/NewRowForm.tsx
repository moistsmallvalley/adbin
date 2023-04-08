import {
  Button,
  Container,
  FormControl,
  FormControlLabel,
  FormLabel,
  Grid,
  Input,
  InputLabel,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
  TextField,
} from "@mui/material";
import { Box } from "@mui/system";
import { skipToken } from "@reduxjs/toolkit/dist/query";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Column, Row } from "./services/payloads";
import { parseColumnValue } from "./services/row";
import {
  useGetRowQuery,
  useGetRowsQuery,
  useGetTableQuery,
  usePatchRowMutation,
  usePostRowMutation,
} from "./services/tables";

export function NewRowForm() {
  const { tableName } = useParams();
  const { data, error, isLoading } = useGetTableQuery(
    tableName ? { tableName } : skipToken
  );
  const [postRow, { isLoading: isPosting }] = usePostRowMutation();
  const [postData, setPostData] = useState<Row>({});

  const create = () => {
    if (!tableName) {
      return;
    }
    postRow({
      tableName: tableName,
      row: postData,
    });
  };

  return error ? (
    <>Network Error</>
  ) : isLoading ? (
    <>Loading...</>
  ) : data ? (
    <Stack width={500}>
      {data.columns.map((c) => (
        <TextField
          key={c.name}
          label={c.name}
          margin="normal"
          fullWidth
          onChange={(e) => {
            setPostData((data) => {
              const newData = structuredClone(data);
              newData[c.name] = parseColumnValue(c, e.target.value);
              return newData;
            });
          }}
        />
      ))}
      <Grid container sx={{ my: 2 }} justifyContent="flex-end">
        <Button variant="contained" onClick={create}>
          Create
        </Button>
      </Grid>
    </Stack>
  ) : null;
}
