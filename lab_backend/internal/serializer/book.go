package serializer

import "time"

type BookingRequest struct {
	BookClass    int       `json:"book_class"`
	BookReason   string    `json:"book_reason"`
	BookTime     time.Time `json:"book_time"` // 可能需要解析为time.Time类型
	BookUsername string    `json:"book_username"`
	LabID        int       `json:"lab_id"`
}
