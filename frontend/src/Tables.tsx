import {
  Box,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import TableRowsOutlinedIcon from "@mui/icons-material/TableRowsOutlined";
import { useGetTablesQuery } from "./services/tables";
import { NavLink } from "react-router-dom";
import { Fragment } from "react";
import { rowsPath } from "./route";

export function Tables() {
  const { data, error, isLoading } = useGetTablesQuery();

  return (
    <Box
      sx={{
        width: "100%",
        maxWidth: 240,
        minHeight: "90vh",
      }}
    >
      {error ? (
        <Fragment>Network Error</Fragment>
      ) : isLoading ? (
        <Fragment>Loading...</Fragment>
      ) : data ? (
        <List>
          {data.map((name) => (
            <ListItem key={name} disablePadding>
              <ListItemButton component={NavLink} to={rowsPath(name)}>
                <ListItemIcon>
                  <TableRowsOutlinedIcon />
                </ListItemIcon>
                <ListItemText primary={name} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      ) : null}
    </Box>
  );
}
