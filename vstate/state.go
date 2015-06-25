


package vstate

type State struct {
	Nodes map[string]string

}

// Creates a new kinda state store
func NewVState() *State {
	s := new(State)
	s.Nodes = make(map[string]string)
	return s

}

// add a node if it dont exist already
func (me *State) AddNode( node string) {
	_, found := me.Nodes[node]
	if found == false {
		me.Nodes[node] = ""
	}
}

// add a load of nodes
func (me *State) AddNodes( nodes []string) {
	for _, n := range nodes {
		me.AddNode(n)
	}
}

// Returns a slice/list of nodes
func (me *State) GetNodes( ) []string {
	lst := make([]string, 0)
	for n, _ := range me.Nodes {
		lst = append(lst, n)
	}
	return lst
}

func (me *State) Update( node, val string ) {
	me.Nodes[node] = val
}

func (me *State) GetNodeVal( node string ) string {
	return me.Nodes[node]
}
