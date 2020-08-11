package models

type CLA struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name" required:"true"`
	Text      string  `json:"text" required:"true"`
	Language  string  `json:"language" required:"true"`
	Submitter string  `json:"submitter" required:"true"`
	Fields    []Field `json:"fields,omitempty"`
}

type Field struct {
	Title       string `json:"title" required:"true"`
	Type        string `json:"type" required:"true"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required" required:"true"`
}

func (this *CLA) Create() error {
	v, err := db.CreateCLA(*this)
	if err == nil {
		this.ID = v
	}

	return err
}

func (this *CLA) Get() error {
	v, err := db.GetCLA(this.ID)
	if err == nil {
		*this = v
	}
	return err
}

func (this *CLA) Delete() error {
	return db.DeleteCLA(this.ID)
}

type CLAs struct {
	BelongTo []string
}

func (this CLAs) Get() ([]CLA, error) {
	return db.ListCLA(this.BelongTo)
}
