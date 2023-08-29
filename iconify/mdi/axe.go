package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2l10 6c0 4-2 6-6 7l-3-5l-4-4l3-4M4.11 19.84l-1.99-1.51L9.19 9L11 10.81l-6.89 9.03Z"/>`),
		g.Group(children),
	)
}