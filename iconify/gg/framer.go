package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Framer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="M12 21V9H6v6l6 6Z"/><path d="M18 9V3H6l6 6H6v6h12l-6-6h6Z"/></g>`),
		g.Group(children),
	)
}