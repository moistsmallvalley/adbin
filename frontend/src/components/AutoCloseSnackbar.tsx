import { Alert, AlertColor, Snackbar } from "@mui/material";
import { useEffect, useState } from "react";

interface AutoCloseSnackbarProps {
  openTrigger: boolean;
  message: string;
  severity: AlertColor;
}

export function AutoCloseSnackbar(props: AutoCloseSnackbarProps) {
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (props.openTrigger) {
      setOpen(true);
    }
  }, [props.openTrigger]);

  return (
    <Snackbar
      open={open}
      anchorOrigin={{ horizontal: "center", vertical: "bottom" }}
      autoHideDuration={2500}
      onClose={() => setOpen(false)}
    >
      <Alert severity={props.severity} sx={{ width: "100%" }}>
        {props.message}
      </Alert>
    </Snackbar>
  );
}
