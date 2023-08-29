package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3H6.25A3.25 3.25 0 0 0 3 6.25V8h5V3Zm0 6.5H3v5h5v-5Zm1.5 5v-5h5v5h-5ZM8 16H3v1.75A3.25 3.25 0 0 0 6.25 21H8v-5Zm1.5 5v-5h5v4.25a.75.75 0 0 1-.75.75H9.5Zm0-13V3h4.25a.75.75 0 0 1 .75.75V8h-5Zm10 12.25a.75.75 0 0 0 1.5 0V3.75a.75.75 0 0 0-1.5 0v16.5Z"/>`),
		g.Group(children),
	)
}