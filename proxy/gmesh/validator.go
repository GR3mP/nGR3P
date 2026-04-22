package gmesh

import (
	"strings"
	"sync"

	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/uuid"
)

type Validator interface {
	Get(id uuid.UUID) *protocol.MemoryUser
	Add(u *protocol.MemoryUser) error
	Del(email string) error
	GetByEmail(email string) *protocol.MemoryUser
	GetAll() []*protocol.MemoryUser
	GetCount() int64
}

func ProcessUUID(id [16]byte) [16]byte {
	id[6] = 0
	id[7] = 0
	return id
}

type MemoryValidator struct {
	email sync.Map
	users sync.Map
}

func (v *MemoryValidator) Add(u *protocol.MemoryUser) error {
	if u.Email != "" {
		_, loaded := v.email.LoadOrStore(strings.ToLower(u.Email), u)
		if loaded {
			return errors.New("User ", u.Email, " already exists.")
		}
	}
	// v.users.Store(ProcessUUID(u.Account.(*MemoryAccount).ID.UUID()), u)
	return nil
}
