package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Utensils(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 8V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v5a4 4 0 0 0 4 4h1a4 4 0 0 0 4-4Zm4 8V2h3a4 4 0 0 1 4 4v10h-4m-3 0v5a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-5M5 12v9a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-9M5 6V2m3 4V2"/>`),
		g.Group(children),
	)
}