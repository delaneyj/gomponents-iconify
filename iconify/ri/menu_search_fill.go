package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuSearchFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.617 13.032a5.5 5.5 0 1 1 1.414-1.414l2.676 2.675l-1.414 1.414l-2.675-2.675ZM3 4h5v2H3V4Zm0 7h5v2H3v-2Zm0 7h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}