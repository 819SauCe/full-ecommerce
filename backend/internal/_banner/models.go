package banner

import "go.mongodb.org/mongo-driver/bson/primitive"

type Banner struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Tittle       string             `json:"tittle"`
	Image        string             `json:"image"`
	Subtittle    string             `json:"subtittle"`
	Button_to    string             `json:"button_to"`
	Button_label string             `json:"button_label"`
}
