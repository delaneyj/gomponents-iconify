package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookArrowClockwiseTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 3.48a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-1 0v.758a4.5 4.5 0 1 0 2 3.742a.5.5 0 0 0-1 0a3.5 3.5 0 1 1-1.696-3H15.5a.5.5 0 0 0 0 1h2ZM16 10.773a5.472 5.472 0 0 1-1 .185V15H5V4a1 1 0 0 1 1-1h3.59c.18-.358.4-.693.651-1H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h9.5a.5.5 0 0 0 0-1H6a1 1 0 0 1-1-1h10a1 1 0 0 0 1-1v-4.227Z"/>`),
		g.Group(children),
	)
}