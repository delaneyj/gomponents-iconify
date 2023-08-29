package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 .698l4 3.334v6.833l5-4.167l7 5.834V22H2V4.032L6 .698ZM10 20h2v-6h6v6h2v-6.532l-5-4.166l-5 4.166V20Zm6 0v-4h-2v4h2Zm-8 0V4.968L6 3.302L4 4.968V20h4Z"/>`),
		g.Group(children),
	)
}