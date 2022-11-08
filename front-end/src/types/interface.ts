export interface MatchedEmails{
	total?     :number  
	hits?: SourceType[]
}

export interface EmailsTypes {
    Body      ?:string
	Date      ?:string
	From      ?:string
	Subject   ?:string
	To        ?:string 
}

export interface SourceType {
	_index     ?:string      
	_type      ?:string      
	_id        ?:string      
	_score     ?:number     
	_source    ?:EmailsTypes  
}