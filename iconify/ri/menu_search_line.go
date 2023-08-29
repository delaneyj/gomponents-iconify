package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuSearchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM10 8.5a5.5 5.5 0 1 1 10.032 3.117l2.675 2.676l-1.414 1.414l-2.675-2.675A5.5 5.5 0 0 1 10 8.5ZM3 4h5v2H3V4Zm0 7h5v2H3v-2Zm18 7v2H3v-2h18Z"/>`),
		g.Group(children),
	)
}