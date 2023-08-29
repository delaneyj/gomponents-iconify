package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 1c.35 0 .687-.06 1-.17V15a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8.17A3 3 0 0 0 16 7Z"/>`),
		g.Group(children),
	)
}