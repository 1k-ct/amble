import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import { createStyles, TextField, Theme } from "@material-ui/core";
import axios, { AxiosError } from "axios";

const useStyles = makeStyles({
  root: {
    minWidth: 275,
    minHeight: 300,
  },
});
const useStyles2 = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      "& .MuiTextField-root": {
        margin: theme.spacing(1),
        width: "60ch",
        outerHeight: 100,
      },
    },
  })
);
type PostTweetItems = {
  // id: number;
  is_private: false;
  name: string;
  content: string;
  like_count: number;
  retweet_count: number;
  reply_count: number;
  // created_at: string;
  // updated_at: string;
};
type PostResponse = {
  msg: string;
};
type IErrorResponse = {
  error: string;
};
export default function TweetCard() {
  const classes = useStyles();
  const classes2 = useStyles2();
  const [value, setValue] = React.useState("");

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.value.length <= 300) {
      setValue(event.target.value);
    }
  };
  const limitCount = (n: number, limit: number): number | null => {
    if (n <= limit) {
      return limit - n;
    }
    return -n + limit;
  };
  const requestData: PostTweetItems = {
    is_private: false,
    name: "ゲスト",
    content: value,
    like_count: 0,
    retweet_count: 0,
    reply_count: 0,
  };
  const postTweet = () => {
    const requestData: PostTweetItems = {
      is_private: false,
      name: "ゲスト",
      content: value,
      like_count: 0,
      retweet_count: 0,
      reply_count: 0,
    };
    axios
      .post<PostResponse>("http://localhost:8080/api/v1/tweet", requestData)
      .then((res) => {
        console.log(res.data.msg);
      })
      .catch((e: AxiosError<IErrorResponse>) => {
        if (e.response !== undefined) {
          console.log(e.response.data.error);
        }
      });
  };
  return (
    <Card className={classes.root}>
      <CardContent>
        <form className={classes2.root} noValidate autoComplete="off">
          <div>
            <TextField
              type="textarea"
              id="outlined-multiline-static"
              label="Tweet"
              multiline
              rows={8}
              rowsMax="15"
              value={value}
              variant="outlined"
              onChange={handleChange}
            />
          </div>
        </form>
      </CardContent>
      <CardActions>
        <Button onClick={postTweet} color="primary" size="medium">
          Tweet
        </Button>
      </CardActions>
      <Typography aria-posinset={10} variant="h5" component="h2">
        {limitCount(value.length, 200)}
      </Typography>
    </Card>
  );
}
