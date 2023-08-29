package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5 5l7 3.5l7-3.5m0 .08v8.42a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1V5.08a1 1 0 0 1 .504-.868l6-3.428a1 1 0 0 1 .992 0l6 3.428a1 1 0 0 1 .504.868Z"/>`),
		g.Group(children),
	)
}