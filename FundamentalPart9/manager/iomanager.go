package manager

type IOManager interface {
	ReadJSON(data interface{}) error
	WriteJSON(data interface{}) error
}
