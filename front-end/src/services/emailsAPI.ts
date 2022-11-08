import { MatchedEmails } from "@/types/interface";
// require("dotenv").config(); // Check

export const emailsSearch = async(term:any): Promise<MatchedEmails> => {
  // const apiURL = process.env.SERVER_URL
  const headers = new Headers()
      headers.append('Access-Control-Allow-Origin', 'http://localhost:3000/api/search');
      headers.append('Content-Type', 'application/json');
      headers.append('Accept', 'application/json');

    const response = await fetch( 'http://localhost:3000/api/search',{
      method: 'POST',
      headers: headers,
      body: JSON.stringify(term)
      })
      
      return await response.json();
}