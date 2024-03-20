package input

type Terminal struct {
	Name string `json:"name"`
}

func (c Terminal) IsValid() error {
	return nil
}
