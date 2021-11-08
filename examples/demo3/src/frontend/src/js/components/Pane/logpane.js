import React from "react";
import { Typography } from "@rmwc/typography";
import { Elevation } from "@rmwc/elevation";
import Grid from "@mui/material/Grid";
import Title from "@/js/components/Title";
import "@material/typography/dist/mdc.typography.css";
export default function LogPane({ data }) {
  // Generate data
  function createData(data) {
    var resp = null;
    if (data) {
      return data.map((element) => {
        return (
          <Typography use="caption" style={{ display: "block" }}>
            {element}
          </Typography>
        );
      });
    }
  }

  return (
    <div>
      <Elevation
        z={2}
        className="component-style "
        style={{ padding: 10, marginBottom: 10 }}
      >
        <Grid
          direction="column"
          container
          style={{ width: "100%", height: "100%" }}
        >
          <Title>Activity Logs</Title>
          <div style={{ height: "45vh", overflowY: "auto" }}>
            {createData(data)}
          </div>
        </Grid>
      </Elevation>
    </div>
  );
}
