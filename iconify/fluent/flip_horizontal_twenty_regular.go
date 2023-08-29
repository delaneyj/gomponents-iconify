package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.921 15.77a.5.5 0 0 1-.421.23h-6a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 .954-.21l6 13a.5.5 0 0 1-.033.48ZM12 4.777V15h4.718L12 4.777ZM2.5 16a.5.5 0 0 1-.454-.71l6-13A.5.5 0 0 1 9 2.5v13a.5.5 0 0 1-.5.5h-6Z"/>`),
		g.Group(children),
	)
}