package skeleton

// Command store command meta information
type Command struct {
	Name string

	// Flags are flag for the command
	Flags []Flag

	// Synopsis is short help message of the command
	Synopsis string

	// Help is long help message of the command
	Help string
}