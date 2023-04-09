import React, { useState } from "react";
import { Button, Grid, Stack, TextField } from "@mui/material";
import { skipToken } from "@reduxjs/toolkit/dist/query";
import { useParams } from "react-router-dom";
import { AutoCloseSnackbar } from "./components/AutoCloseSnackbar";
import { Row } from "./services/payloads";
import { parseColumnValue } from "./services/row";
import { useGetTableQuery, usePostRowMutation } from "./services/tables";

export function NewRowForm() {
  const { tableName } = useParams();
  const { data, error, isLoading } = useGetTableQuery(
    tableName ? { tableName } : skipToken
  );
  const [postRow, { isLoading: isSaving, isSuccess, isError }] =
    usePostRowMutation();
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
        <Button variant="contained" onClick={create} disabled={isSaving}>
          Create
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
