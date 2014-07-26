
package gate

import (
    "time"
)

type User struct {
    Id int64 
    CreatedAt time.Time
    UpdatedAt time.Time  
    Generation int64 
    UserID string `sql:"size:255"`
    UserSecret string `sql:"size:255"`
    QRCode QRCode 
    ScratchCodes []ScratchCode
}

type ScratchCode struct {
    Id int64
    UserId int64 
    CreatedAt time.Time 
    Code string `sql:"size:255"`
}

func (s ScratchCode) String() (string) {
    return s.Code 
}

type QRCode struct {
    Id int64 
    UserId int64 
    CreatedAt time.Time 
    Base64 string 
}