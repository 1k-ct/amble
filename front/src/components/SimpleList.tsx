import React from "react";
import { createStyles, Theme, makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem, { ListItemProps } from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import { indigo, grey } from "@material-ui/core/colors";
import HomeIcon from "@material-ui/icons/Home";
import ExploreIcon from "@material-ui/icons/Explore";
import NotificationsIcon from "@material-ui/icons/Notifications";
import EmailIcon from "@material-ui/icons/Email";
import BookmarkIcon from "@material-ui/icons/Bookmark";
import ListIcon from "@material-ui/icons/List";
import PersonIcon from "@material-ui/icons/Person";
import MoreHorizIcon from "@material-ui/icons/MoreHoriz";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: "100%",
      //   maxWidth: 360,
      backgroundColor: indigo[800],
      color: "#FFFFFF",
    },
    palette: {
      color: grey[300],
    },
  })
);

const SimpleList: React.FC = () => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <List
        className={classes.palette}
        component="nav"
        aria-label="main mailbox folders"
      >
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <HomeIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Home" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <ExploreIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Explore" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <NotificationsIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Notifications" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <EmailIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Messages" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <BookmarkIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Bookmarks" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <ListIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Lists" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <PersonIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="Profile" />
        </ListItem>
        <ListItem button component="a" href="#">
          <ListItemIcon>
            <MoreHorizIcon className={classes.palette} />
          </ListItemIcon>
          <ListItemText primary="More" />
        </ListItem>
      </List>
    </div>
  );
};
export default SimpleList;
