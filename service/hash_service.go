package service

import (
	"errors"
	"hash/crc64"
	"math/bits"
	"time"

	"github.com/gofrs/uuid"

	"github.com/madxiii/tsarka_task/repository"
)

var errWrongHash = errors.New("error: hash length must be 36")

type Hasher struct {
	repo repository.Hash
}

func NewHasher(repo repository.Hash) *Hasher {
	return &Hasher{repo: repo}
}

func (h *Hasher) CalculateBody(key string) (string, error) {
	hash, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	go h.worker(hash, key)

	return hash.String(), nil
}

func (h *Hasher) worker(hash uuid.UUID, key string) {
	var res uint64
	done := make(chan struct{})
	ticker := time.NewTicker(5 * time.Second)

	h.repo.StoreKey(hash.String())

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				res = calc(key)
			}
		}
	}()

	time.Sleep(60 * time.Second)

	ticker.Stop()

	done <- struct{}{}

	ones := countOne(res)

	h.repo.StoreValByKey(hash.String(), ones)
}

func calc(input string) uint64 {
	crc := crc64.Checksum([]byte(input), crc64.MakeTable(crc64.ECMA))

	nanoTime := uint64(time.Now().UnixNano())

	logicAnd := nanoTime & crc

	return logicAnd
}

func countOne(res uint64) int {
	return bits.OnesCount(uint(res))
}

func (h *Hasher) GetResult(param string) (int, error) {
	if len(param) != 36 {
		return 0, errWrongHash
	}

	res, err := h.repo.GetValueByKey(param)
	if err != nil {
		return 0, err
	}

	return res, nil
}
