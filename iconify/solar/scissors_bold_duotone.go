package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6.654 1.633a.75.75 0 1 0-1.308.735L15.704 20.79a3.75 3.75 0 1 0-.136-3.303L6.654 1.633Z" opacity=".5"/><path d="M17.346 1.633a.75.75 0 0 1 1.308.735L8.296 20.79a3.75 3.75 0 1 1 .136-3.303l8.914-15.854Z"/></g>`),
		g.Group(children),
	)
}