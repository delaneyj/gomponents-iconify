package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.25 0l-6 6l-1.5 10.5l7.5 7.5l9-6l-6-6l4.5-4.5l-7.5-7.5Z"/>`),
		g.Group(children),
	)
}