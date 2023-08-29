package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 16.2V7.8c1.2-.4 2-1.5 2-2.8c0-1.7-1.3-3-3-3c-1.3 0-2.4.8-2.8 2H7.8C7.4 2.8 6.3 2 5 2C3.3 2 2 3.3 2 5c0 1.3.8 2.4 2 2.8v8.4c-1.2.4-2 1.5-2 2.8c0 1.7 1.3 3 3 3c1.3 0 2.4-.8 2.8-2h8.4c.4 1.2 1.5 2 2.8 2c1.7 0 3-1.3 3-3c0-1.3-.8-2.4-2-2.8zM16.2 18H7.8c-.3-.8-1-1.5-1.8-1.8V7.8c.8-.3 1.5-1 1.8-1.8h8.4c.3.8 1 1.5 1.8 1.8v8.4c-.8.3-1.5 1-1.8 1.8zM19 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zM5 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm0 16c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1zm14 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1z"/>`),
		g.Group(children),
	)
}