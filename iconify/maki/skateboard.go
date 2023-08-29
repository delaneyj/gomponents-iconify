package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skateboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.25 5.25V5h1.5v.25c0 .69.56 1.25 1.25 1.25h7c.69 0 1.25-.56 1.25-1.25V5h1.5v.25A2.75 2.75 0 0 1 11 8H4a2.75 2.75 0 0 1-2.75-2.75ZM5 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm7 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}