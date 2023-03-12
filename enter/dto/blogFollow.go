package dto

import "go-gin-shop/enter/tb"

type BlogFollow struct {
	List    []tb.TbBlog `json:"list,omitempty"`
	MinTime float64     `json:"minTime,omitempty"`
	Offset  int64       `json:"offset,omitempty"`
}
