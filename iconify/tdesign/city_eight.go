package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h12v8h8v12H2V2Zm10 8V8H9v2h3Zm-3 2v2h3v-2H9Zm-2-2V8H4v2h3Zm-3 2v2h3v-2H4Zm0 4v4h8v-4H4Zm10 4h2v-4h2v4h2v-8h-6v8ZM4 6h3V4H4v2Zm5-2v2h3V4H9Z"/>`),
		g.Group(children),
	)
}