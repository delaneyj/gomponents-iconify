package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.9 19.3l-9-15.6l-.3-.3c-.5-.3-1.1-.2-1.4.3l-9 15.6c-.2.1-.2.3-.2.5c0 .6.4 1 1 1h18c.2 0 .3 0 .5-.1c.5-.3.6-.9.4-1.4z"/>`),
		g.Group(children),
	)
}