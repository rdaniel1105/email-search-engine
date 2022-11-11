import { MatchedEmails, TermType } from "@/types/interface";

let apiURL = process.env.VUE_APP_SERVER_URL;

export const emailsSearch = async(term: TermType,limit: number,from:number) : Promise<MatchedEmails | string> => {
  const headers = new Headers()
  if (apiURL === undefined) {
    apiURL = "http://localhost:3000/api/search"
  }

  headers.append('Access-Control-Allow-Origin',apiURL );
  headers.append('Content-Type', 'application/json');
  headers.append('Accept', 'application/json');

    const response = await fetch(apiURL+`?limit=${limit}&from=${from}`, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(term)
    })

    if (response.status == 500){
      return ("Sorry, we could not find any emails with this term :(");
    }

    return response.json()
}