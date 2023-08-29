package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontalTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.75 14a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0Zm8 0a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0ZM22 16.75a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Z"/>`),
		g.Group(children),
	)
}