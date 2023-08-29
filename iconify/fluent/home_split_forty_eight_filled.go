package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSplitFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M26.372 5.846a3.75 3.75 0 0 0-4.744 0L8.378 16.67A3.75 3.75 0 0 0 7 19.574v19.674a3.75 3.75 0 0 0 3.75 3.75h26.5a3.75 3.75 0 0 0 3.75-3.75V19.574a3.75 3.75 0 0 0-1.377-2.904L26.373 5.846ZM25.25 13.25v2.5a1.25 1.25 0 1 1-2.5 0v-2.5a1.25 1.25 0 1 1 2.5 0ZM24 22c.69 0 1.25.56 1.25 1.25v2.5a1.25 1.25 0 1 1-2.5 0v-2.5c0-.69.56-1.25 1.25-1.25Zm1.25 11.25v2.5a1.25 1.25 0 1 1-2.5 0v-2.5a1.25 1.25 0 1 1 2.5 0Z"/>`),
		g.Group(children),
	)
}