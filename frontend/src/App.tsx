import { Box, Divider, Stack } from "@mui/material";
import { Container } from "@mui/system";
import { useState } from "react";
import { Outlet } from "react-router-dom";
import reactLogo from "./assets/react.svg";
import { Tables } from "./Tables";
import { TableView } from "./TableView";
import viteLogo from "/vite.svg";

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
