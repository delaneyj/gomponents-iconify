package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokeDetectorAlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3H3c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2m0 16H3V5h14v14m-7-1c3.3 0 6-2.7 6-6s-2.7-6-6-6s-6 2.7-6 6s2.7 6 6 6m0-10c2.2 0 4 1.8 4 4s-1.8 4-4 4s-4-1.8-4-4s1.8-4 4-4m13-1h-2v6h2V8m0 7h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}