package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scales(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 704q0 18-15 27q-20 17-59 27t-86 10t-86-10t-59-27q-15-10-15-27q0-6 1-8l116-568H576v769q111 5 183.5 22.5T832 960q0 26-18.5 45t-45.5 19H256q-27 0-45.5-19T192 960q0-23 72.5-40.5T448 897V128H203l116 568q1 2 1 8q0 18-15 27q-20 17-59 27t-86 10t-86-10t-59-27Q0 721 0 704q0-7 1-8L125 88q7-22 32-23l1-.5l2-.5h288q0-27 18.5-45.5T512 0t45.5 18.5T576 64h288q1 0 3 1q25 1 32 23l124 608q1 2 1 8zM79 649q38-9 81-9t81 9l-81-396zm785-396l-81 396q38-9 81-9t81 9z"/>`),
		g.Group(children),
	)
}