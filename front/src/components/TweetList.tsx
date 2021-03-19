import React, { useEffect, useState } from "react";
import { createStyles, Theme, makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemAvatar from "@material-ui/core/ListItemAvatar";
import Avatar from "@material-ui/core/Avatar";
import Typography from "@material-ui/core/Typography";
// import { indigo, grey } from "@material-ui/core/colors";
// import { type } from "node:os";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: "100%",
      //   maxWidth: "36ch",
      //   backgroundColor: theme.palette.background.default,

      // backgroundColor: indigo[500],
      // color: grey[100],
    },
    inline: {
      display: "inline",
    },
  })
);
type Props = {};
type Item = {
  userId: number;
  iconName: string;
  iconSrc: string;
  name: string;
  message: string;
};

const TweetListItem: React.FC<Item> = ({
  userId,
  iconName,
  iconSrc,
  name,
  message,
}) => {
  // const classes = useStyles();
  return (
    <div>
      <ListItem alignItems="flex-start">
        <ListItemAvatar>
          <Avatar alt={iconName} src={iconSrc} />
        </ListItemAvatar>
        <ListItemText
          key={userId}
          primary={name}
          secondary={
            <React.Fragment>
              <Typography component="span" variant="body2" align="left">
                {message}
              </Typography>
            </React.Fragment>
          }
        />
      </ListItem>
      <Divider component="li" />
    </div>
  );
};
const TweetList: React.FC<Props> = (props) => {
  const classes = useStyles();

  const [items, setItems] = useState<Item[]>([]);

  const anyItems: Item[] = [
    {
      userId: 0,
      iconName: "Remy Sharp",
      iconSrc: "static/...jpg",
      name: "Ali Connors",
      message:
        "— I'll be in your neighborhood doing  errands…ddddd dddddddddddd ddddddddddddddddddddddddddddddddd dddddd dddddddddddddddddddddddd dddddddddd dddddddd ddddddddddddddddddddddddddddddddddddddd",
    },
    {
      userId: 1,
      iconName: "Travis Howard",
      iconSrc: "static/...jpg",
      name: "Summer BBQ",
      message: " — Wish I could come, but I'm out of town this…",
    },
    {
      userId: 2,
      iconName: "Cindy Baker",
      iconSrc: "static/...jpg",
      name: "Oui Oui",
      message:
        " — Do you have Paris recommendations? Have you ever…  — Do you have Paris recommendations? Have you ever…  — Do you have Paris recommendations? Have you ever…  — Do you have Paris recommendations? Have you ever…",
    },
    {
      userId: 3,
      iconName: "Cindy Baker",
      iconSrc: "static/...jpg",
      name: "Oui Oui",
      message:
        " — Do you have Paris recommendations? Have you ever…  — Do you have Paris recommendations? Have you ever…  — Do you have Paris recommendations? Have you ever…  — Do you have Paris recommendations? Have you ever…",
    },
  ];

  useEffect(() => {
    setItems(anyItems);
    // console.log(anyItems.map((item) => ({ item })));
    // TODO
  }, [props]);

  return (
    <List className={classes.root}>
      {items.map((i) => (
        <TweetListItem
          key={i.userId}
          userId={i.userId}
          iconName={i.iconName}
          iconSrc={i.iconSrc}
          name={i.name}
          message={i.message}
        />
      ))}
    </List>
  );
};
export default TweetList;
