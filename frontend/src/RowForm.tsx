import React, { FormEventHandler, useEffect, useState } from "react";
import { Button, Grid, Stack, TextField } from "@mui/material";
import { skipToken } from "@reduxjs/toolkit/dist/query";
import { useParams } from "react-router-dom";
import { AutoCloseSnackbar } from "./components/AutoCloseSnackbar";
import { Row } from "./services/payloads";
import { parseColumnValue } from "./services/row";
import { useGetRowQuery, usePatchRowMutation } from "./services/tables";

export function RowForm() {
  const { tableName, keys } = useParams();
  const { data, error, isLoading } = useGetRowQuery(
    tableName && keys ? { tableName, primaryKeys: keys.split("/") } : skipToken
  );
  const [patchRow, { isLoading: isSaving, isSuccess, isError }] =
    usePatchRowMutation();
  const [patchData, setPatchData] = useState<Row>({});

  useEffect(() => {
    setPatchData(data ? structuredClone(data.row) : {});
  }, [data]);

  const save: FormEventHandler = (e) => {
    e.preventDefault();
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
    <Stack component="form" width={500} onSubmit={save}>
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
        <Button variant="contained" type="submit" disabled={isSaving}>
          Save
        </Button>
        <AutoCloseSnackbar
          openTrigger={!isSaving && (isSuccess || isError)}
          message={isSuccess ? "Saved!" : "Error!"}
          severity={isSuccess ? "success" : "error"}
        />
      </Grid>
    </Stack>
  ) : null;
}
