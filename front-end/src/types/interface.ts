export interface MatchedEmails{
	total?     :number  
	hits?: SourceType[]
}

export interface EmailsTypes {
    body      :string
	date      :string
	from      :string
	messageID :string
	subject   :string
	to        :string 
}

export interface SourceType {
	Index     :string      
	Type      :string      
	ID        :string      
	Score     :number     
	TimeStamp :string
	Source    :EmailsTypes   
}