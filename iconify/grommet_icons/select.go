package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Select(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 1h6h-6Zm11.188 18.472L16 22l-3.5-4.5l-3 3.5L7 7l13 6.5l-4.5 1.5l3.688 4.472ZM19 4V1h-3M6 1H3v3m0 10v3h3M19 6v4v-4ZM3 12V6v6Z"/>`),
		g.Group(children),
	)
}