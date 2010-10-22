package main

import ( 
    "log"
)

type Proc interface {
    DoOne(*Handle)
    DoTwo(*Handle)
}

type Handle struct {
    proc Proc
}
func (h *Handle) Run(f string) {
    p := new(h.proc.(type))

    h.proc.DoOne(h)
    h.proc.DoTwo(h)
}

type ProcBase struct {
    Inst int
    Env string
}
func (p *ProcBase) DoOne(h *Handle) {
    log.Printf("ProcBase: DoOne. field=%s, env=%s", p.Env)
}
func (p *ProcBase) DoTwo(h *Handle) {
    log.Printf("ProcBase: DoTwo. field=%s, env=%s", p.Env)
}

type ProcChild struct {
    ProcBase
}
func (p *ProcChild) DoTwo(h *Handle) {
    log.Printf("ProcChild: DoTwo. field=%s, env=%s", p.Env)
}

func NewHandle(p Proc) {
    
}

func main () {

}