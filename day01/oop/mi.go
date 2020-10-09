package oop

// Camera is a struct
type Camera struct{}

// TakePicture is a func
func (_ Camera) TakePicture() string {
	return "Click"
}

// SwitchOn is a func
func (_ Camera) SwitchOn() string {
	return "SwitchOn Camera"
}

// Phone is a struct
type Phone struct{}

// Call is a func
func (_ Phone) Call() string {
	return "Ring Ring"
}

// SwitchOn is a func
func (_ Phone) SwitchOn() string {
	return "SwitchOn Phone"
}

// CameraPhone is a struct
type CameraPhone struct {
	Camera //has anonymous camera
	Phone  //has anonymous phone
}
