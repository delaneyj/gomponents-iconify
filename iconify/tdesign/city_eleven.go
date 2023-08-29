package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2H2v20h20v-9.72l-7-2.334l-1 .333V2Zm-2 8.946l-4 1.333V20H4V4h8v6.946ZM10 20v-6.28l5-1.666l5 1.667V20h-4v-4h-2v4h-4Z"/>`),
		g.Group(children),
	)
}