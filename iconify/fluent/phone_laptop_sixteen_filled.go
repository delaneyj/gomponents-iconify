package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneLaptopSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4a1.5 1.5 0 0 1 1.5-1.5h8A1.5 1.5 0 0 1 14 4v5a1.5 1.5 0 0 1-1.5 1.5h-4v-4A2.5 2.5 0 0 0 6 4H3.5c-.173 0-.34.022-.5.063V4Zm11.5 8.5h-6v-1h6a.5.5 0 0 1 0 1ZM3 5a1.5 1.5 0 0 0-1.5 1.5v6A1.5 1.5 0 0 0 3 14h3a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 6 5H3Zm1.5 7.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}