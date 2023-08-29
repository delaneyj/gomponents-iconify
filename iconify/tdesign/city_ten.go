package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 .72V10h8v12H2V3.72l12-3ZM14 20h2v-5h2v5h2v-8h-6v8ZM4 5.28V20h8V3.28l-8 2Zm2 2.718h2.004v2.004H6V7.998Z"/>`),
		g.Group(children),
	)
}