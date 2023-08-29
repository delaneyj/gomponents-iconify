package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 2C2.343 2 1 3.327 1 4.964v7.048c0 .546.448.988 1 .988h1V8.554c0-2.456 2.015-4.446 4.5-4.446S12 6.098 12 8.554V13h1c.552 0 1-.442 1-.988V4.964C14 3.327 12.657 2 11 2H4Zm7 8v-.952c0-2.183-1.567-3.952-3.5-3.952S4 6.866 4 9.048V13l7-3Z"/>`),
		g.Group(children),
	)
}