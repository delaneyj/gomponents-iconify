package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func City(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 1.5l8 3.333V10h12v12H2V1.5ZM10 12v8h2v-5h6v5h2v-8H10Zm6 8v-3h-2v3h2Zm-8 0V6.167L4 4.5V20h4Z"/>`),
		g.Group(children),
	)
}