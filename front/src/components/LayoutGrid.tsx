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
    padding: theme.spacing(2),
    textAlign: "center",
    color: theme.palette.text.secondary,
  },
}));

const CenteredGrid: React.FC = () => {
  const classes = useStyles();

  return (
    <Container>
      <div className={classes.root}>
        <Grid container>
          <Grid item xs={2}>
            <SimpleList />
            <Typography
              component="div"
              style={{ backgroundColor: "#cfe8fc", height: "100%" }}
            />
          </Grid>
          <Grid item xs={7}>
            {/* <Paper className={classes.paper}>xs=6</Paper> */}
            <TweetList />
          </Grid>
          <Grid item xs={3}>
            <Paper className={classes.paper}>xs=3</Paper>
          </Grid>
        </Grid>
      </div>
    </Container>
  );
};
export default CenteredGrid;
