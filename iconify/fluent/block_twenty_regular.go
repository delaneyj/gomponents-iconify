package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16zm0 1a7 7 0 1 0 0 14a7 7 0 0 0 0-14zM6.5 9.5h7a.5.5 0 0 1 .09.992l-.09.008h-7a.5.5 0 0 1-.09-.992L6.5 9.5h7h-7z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}