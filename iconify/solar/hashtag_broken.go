package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashtagBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 3L8.75 7.5l-.139.5l-.833 3M5 21l1.944-7M19 3l-5 18m8-12H4m16 7H2"/>`),
		g.Group(children),
	)
}