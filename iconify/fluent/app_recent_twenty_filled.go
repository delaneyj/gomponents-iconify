package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppRecentTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H7ZM2 6a2 2 0 0 1 2-2v12a2 2 0 0 1-2-2V6Zm14 10V4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}