package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9v1.5l-6 6l-6-6V9h12m-6 4.67L14.67 11H9.33L12 13.67Z"/>`),
		g.Group(children),
	)
}