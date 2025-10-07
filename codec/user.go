package codec

type User struct {
	Name  string `codec:"name"`
	Email string `codec:",omitempty"`
}
