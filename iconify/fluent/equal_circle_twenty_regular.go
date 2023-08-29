package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualCircleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 9a.5.5 0 0 0 0-1h-7a.5.5 0 0 0 0 1h7Zm0 3a.5.5 0 0 0 0-1h-7a.5.5 0 0 0 0 1h7Zm4.5-2a8 8 0 1 0-16 0a8 8 0 0 0 16 0Zm-8-7a7 7 0 1 1 0 14a7 7 0 0 1 0-14Z"/>`),
		g.Group(children),
	)
}