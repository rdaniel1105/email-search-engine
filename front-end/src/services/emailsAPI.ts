import type { SearchResponse, SearchTerm } from '@/types/interface';

const DEFAULT_API_URL = 'http://localhost:3000/api/v1/emails/search';

const apiURL = (process.env.VUE_APP_SERVER_URL?.trim() || DEFAULT_API_URL).replace(/\/+$/, '');

export class SearchError extends Error {
  constructor(message: string, public readonly status?: number) {
    super(message);
    this.name = 'SearchError';
  }
}

interface SearchParams {
  term: SearchTerm;
  limit: number;
  from: number;
  signal?: AbortSignal;
}

export async function emailsSearch({ term, limit, from, signal }: SearchParams): Promise<SearchResponse> {
  const url = `${apiURL}?limit=${encodeURIComponent(limit)}&from=${encodeURIComponent(from)}`;

  let response: Response;

  try {
    response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify(term),
      signal,
    });
  } catch (cause) {
    if ((cause as DOMException)?.name === 'AbortError') {
      throw cause;
    }
    throw new SearchError('The archive is unreachable. Check the server connection.', 0);
  }

  if (!response.ok) {
    if (response.status === 400) {
      throw new SearchError('The query was rejected. Please refine your search.', 400);
    }
    if (response.status === 413) {
      throw new SearchError('The query is too large.', 413);
    }
    throw new SearchError('The archive responded with an unexpected error.', response.status);
  }

  const data = (await response.json()) as SearchResponse;
  return data;
}
