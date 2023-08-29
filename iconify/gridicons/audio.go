package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4v10.184A2.962 2.962 0 0 0 7 14a3 3 0 1 0 3 3V7h7v4.184A2.962 2.962 0 0 0 16 11a3 3 0 1 0 3 3V4H8z"/>`),
		g.Group(children),
	)
}