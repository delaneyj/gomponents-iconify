package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v20H4V2Zm2 2v4h12V4H6Zm12 6h-3v2h3v-2Zm0 4h-3v2h3v-2Zm0 4h-3v2h3v-2Zm-5 2v-2h-2v2h2Zm-4 0v-2H6v2h3Zm-3-4h3v-2H6v2Zm0-4h3v-2H6v2Zm5-2v2h2v-2h-2Zm2 4h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}