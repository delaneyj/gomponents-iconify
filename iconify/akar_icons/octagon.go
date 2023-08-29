package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11.235 2.317a2 2 0 0 1 1.53 0l5.54 2.295a2 2 0 0 1 1.083 1.082l2.295 5.54a2 2 0 0 1 0 1.531l-2.295 5.54a2 2 0 0 1-1.082 1.083l-5.54 2.295a2 2 0 0 1-1.531 0l-5.54-2.295a2 2 0 0 1-1.083-1.082l-2.295-5.54a2 2 0 0 1 0-1.531l2.295-5.54a2 2 0 0 1 1.082-1.083l5.54-2.295Z"/>`),
		g.Group(children),
	)
}