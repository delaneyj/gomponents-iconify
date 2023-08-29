package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5 14.5l8-8m-6 6l2 2m0-4l2 2m4.5-5a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Z"/>`),
		g.Group(children),
	)
}