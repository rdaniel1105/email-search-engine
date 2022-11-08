import { MatchedEmails } from "@/types/interface";

export const emailsSearch = async(term:any): Promise<MatchedEmails> => {

  const headers = new Headers()
      headers.append('Access-Control-Allow-Origin', 'http://localhost:3000/api/search');
      headers.append('Content-Type', 'application/json');
      headers.append('Accept', 'application/json');

    const response = await fetch( 'http://localhost:3000/api/search',{
      method: 'POST',
      headers: headers,
      body: JSON.stringify(term)
      })
      
      return response.json();
    }