package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontalTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.25 14a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0Zm8 0a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0ZM22 16.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5Z"/>`),
		g.Group(children),
	)
}