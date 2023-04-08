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
  usePatchRowMutation,
} from "./services/tables";

export function RowForm() {
  const { tableName, keys } = useParams();
  const { data, error, isLoading } = useGetRowQuery(
    tableName && keys ? { tableName, primaryKeys: keys.split("/") } : skipToken
  );
  const [patchRow, { isLoading: isPatching }] = usePatchRowMutation();
  const [patchData, setPatchData] = useState<Row>({});

  useEffect(() => {
    setPatchData(data ? structuredClone(data.row) : {});
  }, [data]);

  const save = () => {
    if (!tableName || !keys) {
      return;
    }
    patchRow({
      tableName: tableName,
      primaryKeys: keys.split("/"),
      row: patchData,
    });
  };

  return error ? (
    <>Network Error </>
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
          defaultValue={data.row[c.name] ?? ""}
          onChange={(e) => {
            setPatchData((data) => {
              const newData = structuredClone(data);
              newData[c.name] = parseColumnValue(c, e.target.value);
              return newData;
            });
          }}
        />
      ))}
      <Grid container sx={{ my: 2 }} justifyContent="flex-end">
        <Button variant="contained" onClick={save}>
          Save
        </Button>
      </Grid>
    </Stack>
  ) : null;
}
