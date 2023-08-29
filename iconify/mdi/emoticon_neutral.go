package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmoticonNeutral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2M7 9.5A1.5 1.5 0 0 1 8.5 8A1.5 1.5 0 0 1 10 9.5A1.5 1.5 0 0 1 8.5 11A1.5 1.5 0 0 1 7 9.5m9 6.5H8v-2h8v2m-.5-5A1.5 1.5 0 0 1 14 9.5A1.5 1.5 0 0 1 15.5 8A1.5 1.5 0 0 1 17 9.5a1.5 1.5 0 0 1-1.5 1.5Z"/>`),
		g.Group(children),
	)
}