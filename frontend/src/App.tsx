import React from "react";
import { Divider, Stack } from "@mui/material";
import { Outlet } from "react-router-dom";
import { Tables } from "./Tables";

function App() {
  return (
    <Stack direction="row" spacing={2} sx={{ p: 2 }}>
      <Tables />
      <Divider orientation="vertical" flexItem />
      <Outlet />
    </Stack>
  );
}

export default App;
