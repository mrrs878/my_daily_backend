package pushSub

type Suber struct {
	Label string
	Cb    func(params interface{})
}
type Pusher struct {
	Label string
}

var subers = make(map[string][]Suber)
var Pushers Pusher

func (suber Suber) Sub(label string) {
	tmp := subers[label]
	if tmp == nil {
		subers[label] = make([]Suber, 0)
	}
	subers[label] = append(subers[label], suber)
}
func (pusher *Pusher) Push(label string, params interface{}) {
	tmp := subers[label]
	if len(tmp) == 0 {
		return
	}
	for _, suber := range tmp {
		suber.Cb(params)
	}
}
