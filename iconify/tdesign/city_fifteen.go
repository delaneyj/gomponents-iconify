package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2H2v20h20V10.687L14 9.02V2Zm-2 6.604L8 7.77V20H4V4h8v4.604ZM10 20v-9.77l10 2.083V20h-4v-4h-2v4h-4Z"/>`),
		g.Group(children),
	)
}