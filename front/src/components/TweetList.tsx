import React, { useEffect, useMemo, useState } from "react";
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
import { GetAPIItems, GetTweets } from "./rest/api";
import axios, { AxiosError } from "axios";

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
type GetItems = {
  id: number;
  is_private: false;
  name: string;
  content: string;
  like_count: number;
  retweet_count: number;
  reply_count: number;
  created_at: string;
  updated_at: string;
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

  type IPostRequest = {
    ids: number[];
  };

  type IPostResponse = {
    id: number;
    is_private: false;
    name: string;
    content: string;
    like_count: number;
    retweet_count: number;
    reply_count: number;
    created_at: string;
    updated_at: string;
  };

  type IErrorResponse = {
    error: string;
  };

  const requestData: IPostRequest = {
    ids: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 18, 19, 20, 21],
  };

  const [tweetItems, setTweetItems] = useState<IPostResponse[]>([]);

  useEffect(() => {
    axios
      .post<IPostResponse[]>("http://localhost:8080/api/v1/tweets", requestData)
      .then((res) => {
        setTweetItems(res.data);
        console.log(res.data);
      })
      .catch((e: AxiosError<IErrorResponse>) => {
        if (e.response !== undefined) {
          console.log(e.response.data.error);
        }
      });
    // TODO
  }, []);

  return (
    <List className={classes.root}>
      {tweetItems.map((i) => (
        <TweetListItem
          key={i.id}
          userId={i.id}
          iconName={i.name}
          iconSrc={i.name}
          name={i.name}
          message={i.content}
        />
      ))}
    </List>
  );
};
export default TweetList;
// {
//   items.map((i) => (
//     <TweetListItem
//       key={i.userId}
//       userId={i.userId}
//       iconName={i.iconName}
//       iconSrc={i.iconSrc}
//       name={i.name}
//       message={i.message}
//     />
//   ));
// }
