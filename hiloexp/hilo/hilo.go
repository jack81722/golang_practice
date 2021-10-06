package hilo

import (
	"errors"
	"sync"
)

type HiLoGen struct {
	mux       sync.Mutex
	share_hi  int64
	max_lo    int64
	incr_size int64
}

func NewHiLoGen(offset, max_lo, incr_size int64) *HiLoGen {
	if max_lo <= 0 {
		panic(errors.New("max_lo must be greater than zero"))
	}
	if incr_size <= 0 {
		panic(errors.New("incr_size must be greater than zero"))
	}
	return &HiLoGen{
		share_hi:  offset,
		max_lo:    max_lo,
		incr_size: incr_size,
	}
}

type HiLoToken struct {
	cur_hi int64
	cur_lo int64
}

func (h *HiLoGen) NewToken() *HiLoToken {
	t := &HiLoToken{}
	h.Next(t)
	return t
}

func (h *HiLoGen) Next(t *HiLoToken) int64 {
	h.mux.Lock()
	hi := h.share_hi
	h.share_hi = h.share_hi + h.incr_size
	t.cur_hi = hi
	t.cur_lo = 0
	h.mux.Unlock()
	return hi
}

func (h *HiLoGen) Key(token *HiLoToken) int64 {
	if token.cur_lo >= h.max_lo {
		h.Next(token)
	}
	key := token.cur_hi*h.max_lo + token.cur_lo
	token.cur_lo = token.cur_lo + 1
	return key
}
