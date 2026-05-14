// Domain types for the search response. Mirrors the back-end's EmailResponse
// (see back-end/models/email.go) and the JSON-encoded wrapper returned from
// /api/v1/emails/search.

export interface EmailSource {
  Body?: string;
  Date?: string;
  From?: string;
  'Message-ID'?: string;
  Subject?: string;
  To?: string;
}

export interface EmailHit {
  _id?: string;
  _index?: string;
  _score?: number;
  _source?: EmailSource;
}

export interface SearchTotal {
  value: number;
}

export interface SearchResponse {
  total?: SearchTotal;
  hits?: EmailHit[];
}

export interface SearchTerm {
  term: string;
}

export type SearchStatus = 'idle' | 'loading' | 'success' | 'empty' | 'error';
