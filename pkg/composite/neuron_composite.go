package composite

// - objects can use either objects via composition, and
// - some composed and singular objects need either similar or in fact, identical behaviors. So composite design pattern lets us treat both types of these objects uniformly, and we can
// support things like iteration.

// - you can iterate composite objects, but you can also iterate just singular objects.

// n1 (out) -> (in) n2 -> n3
// n1 (out) -> n3 -> (in) n2
// n1 has multiple outs
// n2 has multiple ins
type Neuron struct {
	Name    string
	Rank    string
	In, Out []*Neuron
}

func NewNeuron(name string, rank string) *Neuron {
	return &Neuron{Name: name, Rank: rank}
}

func (n *Neuron) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	result = append(result, n)

	return result
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeutronLayer(count int) *NeuronLayer {
	nl := NeuronLayer{}
	nl.Neurons = make([]Neuron, count)
	return &nl
}

func (nl *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)

	for i, _ := range nl.Neurons {
		result = append(result, &nl.Neurons[i])
	}
	return result
}

// composite type
type Composite interface {
	Iter() []*Neuron
}

// Goal: to define a function that works with both at Neuron object & at a collection of Neuron object such as NeuronLayer.
// e.g. Connect(n1, n2)  ; n1->n2
// Connect(n1, {n2,n3})  ; n1->n2, n1->n3
// Connect({n4,n5}, {n6}); n4->n6, n5->n6
// c1 -> c2 composite
func Connect(c1, c2 Composite) {
	for _, u := range c1.Iter() {
		for _, v := range c2.Iter() {
			connectImpl(u, v)
		}
	}
}

func connectImpl(n1 *Neuron, n2 *Neuron) {
	//n1 -> n2
	n1.Out = append(n1.Out, n2)
	n2.In = append(n2.In, n1)
}
