package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharedfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1025H64q-27 0-45.5-18.5T0 961t18.5-45.5T64 897h384V769H256q-27 0-45.5-18.5T192 705V65q0-27 18.5-45.5T256 1h256v288q0 13 9.5 22.5T544 321h287l1 1v383q0 27-19 45.5T768 769H576v128h384q27 0 45.5 19t18.5 45.5t-18.5 45T960 1025zM576 0q26 0 44 18l193 194q19 18 18 45H576V0z"/>`),
		g.Group(children),
	)
}