Oh. I'm open. Speed matters :) But interfaces NOT solve all problems. Ok. Simple production problem: set of handlers for processing.

On first try (Paradigm implicit in the slides from http://golang.org/doc/GoCourseDay2.pdf):

    package main
    
    import ( 
        "log"
    )
    
    type Handler interface {
        Handle() // Calls Do()
        Do()
    }
    
    // Base handler
    type B struct {
        Data string  // Some variable data
    }
    func (b *B) Handle() {
        b.Do()
    }
    func (b *B) Do() {
        log.Print("B: ", b.Data)
    }
    
    // Child
    type C struct {
        B
    }
    func (c *C) Do() {
        log.Print("C1: ", c.Data)
    }
    
    func main () {
        h := &C{B{"data"}} // Why not &C{"data"}?
        h.Handle()
    }

BANG! B.Handle() calls only B.Do() (from language spec)

        2010/10/23 01:58:47 B: data

Second try from http://groups.google.com/group/golang-nuts/browse_thread/thread/7cc11e5324bfc39f/a87127871783769c. 

    package main
    
    import ( 
        "log"
    )
    
    type Handler interface {
        Handle() // Calls Do()
        Do()
    }
    
    // Base handler
    type B struct {
        self Handler
        Data string  // Some variable data
    }
    // Fake constructor
    func (b *B) Init(data string) {
        b.Data = data
    }
    func (b *B) Handle() {
        b.self.Do()
    }
    func (b *B) Do() {
        log.Print("B: ", b.Data)
    }
    
    // Child
    type C struct {
        B
    }
    // Need set self in each child
    func (c *C) Init(data string) {
        c.B.Init(data)  // Call base fake constructor
        c.self = c      // set instance
    }
    func (c *C) Do() {
        log.Print("C1: ", c.Data)
    }
    // Child
    type C1 struct {
        C
    }
    // Need set self in each child
    func (c *C1) Init(data string) {
        c.C.Init(data)  // Call base fake constructor
        c.self = c      // set instance
    }
    func (c *C1) Do() {
        log.Print("C2: ", c.Data)
    }
    
    func main () {
        h := &C{}
        h.Init("data1")
        h.Handle()
        
        h1 := &C1{}
        h1.Init("data2")
        h1.Handle()
    }


Success! 
        
        "2010/10/23 02:08:45 C1: data1"
        "2010/10/23 02:08:45 C2: data2"

Some frustrating code but works OK. It's not inheritance. It's composition. And it's can be done by another (Go) way:

    package main
    
    import ( 
        "log"
    )
    
    type Doer interface {
        Do(*Handle)    // Calls Serv
        Serv()  // Service function
    }
    
    // Handler. Keeps data and runs Doer.Do() 
    type Handle struct {
        Data string
        D Doer
    }
    func (h *Handle) Handle(){
        h.D.Do(h)
    }
    
    // Doer one. Not keeps any data. 
    // Calls other self methods only through Handle.D 
    type Do1 struct {}
    func (d *Do1) Do(h *Handle) {
        log.Print("Do1", h.Data)
        h.D.Serv()
    }
    func (d *Do1) Serv() {
        log.Print("Serv1")
    }
    
    type Do2 struct { Do1 }
    func (d *Do2) Serv() {
        log.Print("Serv2")
    }
    
    func main () {
        h1 := &Handle{"data1", &Do1{}}
        h2 := &Handle{"data2", &Do2{}}
    
        h1.Handle()
        h2.Handle()
    }

Another win:

        2010/10/23 03:03:28 Do1data1
        2010/10/23 03:03:28 Serv1
        2010/10/23 03:03:28 Do1data2
        2010/10/23 03:03:28 Serv2


On 23 îêò, 00:36, Russ Cox <r...@golang.org> wrote:
> > - Normal type interihance. Now for handling different nature of the
> > data that come from the same source I have come up with a strange
> > design with a lot of crutches. Yes I talk about normal classes. Go is
> > not functional language - without normal OO system it useless.
> 
> I think you'll find that people disagree on exactly what OO means.
> It doesn't have to mean type hierarchies and type inheritance.
> Alan Kay famously once said "I invented the term Object-Oriented,
> and I can tell you I did not have C++ in mind."
> 
> I doubt Go will ever have a C++/Java-style type hierarchy.
> The interface idea is one of the core features of the
> language.  If you're not open to learning a new way to
> think about programming, you're probably better off sticking
> with C++ or Java or whatever you're currently using.
> 
> Russ