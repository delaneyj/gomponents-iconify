package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.75A3.75 3.75 0 0 1 6.75 3h6.5v10.25H3v-6.5Zm0 8v6.5A3.75 3.75 0 0 0 6.75 25h6.5V14.75H3ZM14.75 25h6.5A3.75 3.75 0 0 0 25 21.25v-6.5H14.75V25ZM25 13.25v-6.5A3.75 3.75 0 0 0 21.25 3h-6.5v10.25H25Z"/>`),
		g.Group(children),
	)
}