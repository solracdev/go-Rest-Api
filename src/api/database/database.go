package database

// Connect method
type Connect interface {
	Connect() error
}

// Disconnect method
type Disconnect interface {
	Disconnect() error
}
