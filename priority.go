package ice

import "github.com/gortc/stun"

// PriorityAttr represents PRIORITY attribute.
type PriorityAttr uint32

const prioritySize = 4 // 32 bit

// AddTo adds PRIORITY attribute to message.
func (p PriorityAttr) AddTo(m *stun.Message) error {
	v := make([]byte, prioritySize)
	bin.PutUint32(v, uint32(p))
	m.Add(stun.AttrPriority, v)
	return nil
}

// GetFrom decodes PRIORITY attribute from message.
func (p *PriorityAttr) GetFrom(m *stun.Message) error {
	v, err := m.Get(stun.AttrPriority)
	if err != nil {
		return err
	}
	if err = stun.CheckSize(stun.AttrPriority, len(v), prioritySize); err != nil {
		return err
	}
	*p = PriorityAttr(bin.Uint32(v))
	return nil
}
