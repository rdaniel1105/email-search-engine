import { MatchedEmails } from "@/types/interface";
// require("dotenv").config(); // Check
// from: number, limit: number
// const url = new URL(`https://localhost:8080?from=${from}&limit=${limit}`);
// const queryParams = new URLSearchParams(url.search);


export const emailsSearch = async (term: any,limit: number,from:number): Promise<MatchedEmails> => {//TODO error
  // const apiURL = process.env.SERVER_URL
  const headers = new Headers()
  headers.append('Access-Control-Allow-Origin', 'http://localhost:3000/api/search');
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

  const response = await fetch(`http://localhost:3000/api/search?limit=${limit}&from=${from}`, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(term)
  })

  return await response.json();
}