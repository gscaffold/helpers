package rest

import "errors"

// IDRequest 简单封装一个只有 id 参数的请求体.
type IDRequest struct {
	ID int64 `uri:"id" form:"id" json:"id" binding:"required"`
}

func (req *IDRequest) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	return nil
}

type IDRequestUint struct {
	ID uint64 `json:"id"`
}

func (req *IDRequestUint) Validate() error {
	if req.ID == 0 {
		return errors.New("must have id")
	}
	return nil
}

type IDRequestString struct {
	ID string `json:"id"`
}

func (req *IDRequestString) Validate() error {
	if req.ID == "" {
		return errors.New("must have id")
	}
	return nil
}
