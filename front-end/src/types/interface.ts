// Public search response — mirrors the back-end's SearchResponse / SearchHit
// types (see back-end/models/email.go). Flat, snake-case, no leaking of
// ZincSearch's storage fields.

export interface EmailHit {
  id: string;
  score: number;
  from: string;
  to: string;
  subject: string;
  date: string;
  body: string;
}

export interface SearchResponse {
  total: number;
  hits: EmailHit[];
}

export interface SearchTerm {
  term: string;
}

export type SearchStatus = 'idle' | 'loading' | 'success' | 'empty' | 'error';
