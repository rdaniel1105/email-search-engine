export interface MatchedEmails {
	total?: TotalType
	hits?: SourceType[]
}

export interface TotalType {
	value: number
}

export interface EmailsTypes {
	Body?: string
	Date?: string
	From?: string
	Subject?: string
	To?: string
}

export interface SourceType {
	_index?: string
	_type?: string
	_id?: string
	_score?: number
	_source?: EmailsTypes
}

export interface TermType {
	term: string;
  }