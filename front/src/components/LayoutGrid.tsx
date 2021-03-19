import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import SimpleList from "./SimpleList";
import TweetList from "./TweetList";
import Typography from "@material-ui/core/Typography";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    minWidth: 170,
  },
}));

const CenteredGrid: React.FC = () => {
  const classes = useStyles();

  return (
    <Container>
      <div className={classes.root}>
        <Grid container>
          <Grid className={classes.paper} item xs={2}>
            <SimpleList />
            <Typography
              component="div"
              style={{ backgroundColor: "#cfe8fc", height: "100%" }}
            />
          </Grid>
          <Grid item xs={7}>
            <TweetList />
          </Grid>
          <Grid className={classes.paper} item xs={3}>
            <Paper
              style={{ backgroundColor: "#cfe8fc", height: "100%" }}
              className={classes.paper}
            ></Paper>
          </Grid>
        </Grid>
      </div>
    </Container>
  );
};
export default CenteredGrid;
