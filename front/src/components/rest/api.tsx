// import React, {
//   useEffect,
//   createContext,
//   useState,
//   useReducer,
//   useMemo,
// } from "react";
import axios, { AxiosError } from "axios";
type IPostRequest = {
  ids: number[];
};

type IPostResponse = {
  ids: number[];
};

type IErrorResponse = {
  error: string;
};
// const requestData: IPostRequest = {
//   email: "eve.holt@reqres.in",
//   password: "pistol",
// };
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
type ItemDet = {
  id: number;
  name: string;
  year: number;
  color: string;
  pantoneValue: string;
};
export function GetTweets(url: string, requestData: IPostRequest): any {
  axios
    .post<GetItems>(url, requestData)
    .then((res) => {
      return res.data;
    })
    .catch((e: AxiosError<IErrorResponse>) => {
      if (e.response !== undefined) {
        return e.response.data.error;
      }
    });
}
export const GetAPIItems = (url: string): any => {
  axios
    .get<GetItems>(url)
    .then((res) => {
      return res.data;
    })
    .catch((e: AxiosError<IErrorResponse>) => {
      if (e.response !== undefined) {
        // console.log(e.response.data.error);
        return e.response.data.error;
      }
    });
};
