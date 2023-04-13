package models

import "time"

type User struct {
  Email string `bson:"email,omitempty"`
  Password string `bson:"password,omitempty"`
  Created time.Time `bson:"createdAt"`
}
