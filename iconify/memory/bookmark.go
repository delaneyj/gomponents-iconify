package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 2h12v1h1v17h-2v-1h-2v-1h-2v-1h-2v1H8v1H6v1H4V3h1V2m1 2v13h1v-1h2v-1h1v-1h2v1h1v1h2v1h1V4H6Z"/>`),
		g.Group(children),
	)
}