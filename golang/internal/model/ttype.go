package model

import (
	"fmt"
	"time"
)

// A Ttype represents a meaninglessness of our life
type Ttype struct {
	id         int
	cT         string // время создания
	fT         string // время выполнения
	taskRESULT []byte
}

func (t Ttype) String() string {
	return fmt.Sprintf("id=%d timeAdd=%s timeWork=%s result=ok message=%s", t.id, t.cT, t.fT, t.taskRESULT)
}

func (task Ttype) IsTaskFailed(timeout int) bool {
	t, err := time.Parse(time.RFC3339, task.cT)
	return err != nil || time.Now().After(t.Add(time.Second*time.Duration(timeout)))
}

func (t Ttype) GetId() int {
	return t.id
}

func (t *Ttype) SetId(id int) {
	t.id = id
}

func (t Ttype) GetCT() string {
	return t.cT
}

func (t *Ttype) SetCT(cT string) {
	t.cT = cT
}

func (t Ttype) GetFT() string {
	return t.fT
}

func (t *Ttype) SetFT(fT string) {
	t.fT = fT
}

func (t Ttype) GetTaskRESULT() []byte {
	return t.taskRESULT
}
func (t *Ttype) SetTaskRESULT(taskRESULT []byte) {
	t.taskRESULT = taskRESULT
}

func NewTtype(id int, cT string) Ttype {
	return Ttype{id: id, cT: cT}
}
