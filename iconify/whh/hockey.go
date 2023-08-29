package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hockey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m992.57 769l-311-31q-7 1-9-1l-642-642q-22-22-28-50.5t5-36.5q8-11 36.5-5t51.5 28l604 605l293-59q13 0 22.5 9.5t9.5 22.5v128q0 13-9.5 22.5t-22.5 9.5zm-512 128q40 0 68 19t28 45.5t-28 45t-68 18.5t-68-18.5t-28-45t28-45.5t68-19z"/>`),
		g.Group(children),
	)
}