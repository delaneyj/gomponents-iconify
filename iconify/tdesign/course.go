package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Course(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3h-2v6.5l-3-2.25l-3 2.25V3H5v18h14V3Zm-6 0v2.5l1-.75l1 .75V3h-2Zm8 20H3V1h18v22Z"/>`),
		g.Group(children),
	)
}