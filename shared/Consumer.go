package shared

const (
	// Event to pass to Consumer.Control.
	// Will cause the consumer to halt and shutdown.
	ConsumerControlStop = 1

	// Event set by the consumer after Consume() has stopped
	ConsumerControlResponseDone = 1
)

// Consumers are plugins that process messages generated by producer plugins to
// e.g. write them into a logfile.
type Consumer interface {

	// Create a new instance of the concrete plugin class implementing this
	// interface. Expect the instance passed to this function to not be
	// initialized.
	Create(PluginConfig) (Consumer, error)

	// Main loop that fetches messages from a given source and pushes it to the
	// message channel.
	Consume()

	// Returns write access to this consumer's control channel.
	// See ConsumerControl* constants.
	Control() chan<- int

	ControlResponse() <-chan int

	// Returns read access to the message channel this consumer writes to.
	Messages() <-chan Message
}