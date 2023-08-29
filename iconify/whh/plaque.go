package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plaque(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 1024H192q0-80-56.5-136T0 832V192q80 0 136-56T192 0h640q0 80 56 136t136 56v640q-80 0-136 56t-56 136zm64-768q-53 0-90.5-37.5T768 128H256q0 53-37.5 90.5T128 256v512q53 0 90.5 37.5T256 896h512q0-53 37.5-90.5T896 768V256zM704 832H320q-10-47-45.5-82.5T192 704V320q47-10 82.5-45.5T320 192h384q10 47 45.5 82.5T832 320v384q-47 10-82.5 45.5T704 832z"/>`),
		g.Group(children),
	)
}