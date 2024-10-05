package models

type MessageType byte

const (
	// Quote Messages.
	QuoteType MessageType = 'T'
	// Book Messages.
	BookType MessageType = 'B'
	// News Messages.
	SynType MessageType = 'S'
)
