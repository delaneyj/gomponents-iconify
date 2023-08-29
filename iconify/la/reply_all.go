package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m10.281 5.281l-8 8l-.687.719l.687.719l8 8l1.438-1.438L4.438 14l7.28-7.281zm5 0l-8 8l-.687.719l.687.719l8 8l1.438-1.438L10.437 15H23c2.773 0 5 2.227 5 5s-2.227 5-5 5v2c3.855 0 7-3.145 7-7s-3.145-7-7-7H10.437l6.282-6.281z"/>`),
		g.Group(children),
	)
}