package pipe

import (
	"fmt"
	"strings"

	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/services"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/pkg/utils"
	"github.com/Doozers/adapterKitty/AK-interfaces/go-console/proto"
)

type Pipe interface {
	Follow(s []string) (*proto.AdapterRequest, services.GrpcType, error)
	Pipe() *Pipe
	getId() string
}

type PipeWay struct {
	ID     string
	Branch []*Pipe
	Check  *utils.CheckOpt
}

func (p *PipeWay) check(s []string) bool {
	if p.Check == nil {
		return true
	}
	return utils.CheckArgs(s, p.Check)
}

func (p *PipeWay) Follow(a []string) (*proto.AdapterRequest, services.GrpcType, error) {
	if !p.check(a) {
		return nil, 0, fmt.Errorf("invalid arguments")
	}

	for _, p := range p.Branch {
		if (*p).getId() == a[0] {
			return (*p).Follow(a[1:])
		}
	}
	return nil, 0, nil
}

func (p *PipeWay) Pipe() *Pipe {
	P := Pipe(p)
	return &P
}

func (p *PipeWay) getId() string {
	return p.ID
}

type PipeEnd struct {
	ID    string
	F     func([]string) (*proto.AdapterRequest, services.GrpcType, error)
	Check *utils.CheckOpt
}

func (p *PipeEnd) check(s []string) bool {
	if p.Check == nil {
		return true
	}
	return utils.CheckArgs(s, p.Check)
}

func (p *PipeEnd) Follow(s []string) (*proto.AdapterRequest, services.GrpcType, error) {
	if !p.check(s) {
		return nil, 0, nil
	}
	return p.F(s)
}

func (p *PipeEnd) Pipe() *Pipe {
	P := Pipe(p)
	return &P
}

func (p *PipeEnd) getId() string {
	return p.ID
}

type Pipeline struct {
	Pipes []*Pipe
}

func (p *Pipeline) Piping(s string) (*proto.AdapterRequest, services.GrpcType, error) {
	a := strings.Split(s, " ")
	for _, p := range p.Pipes {
		if (*p).getId() == a[0] {
			return (*p).Follow(a[1:])
		}
	}

	return nil, 0, nil
}
