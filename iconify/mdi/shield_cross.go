package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1L3 5v6c0 5.5 3.8 10.7 9 12c5.2-1.3 9-6.5 9-12V5l-9-4m4 9h-3v8h-2v-8H8V8h3V5h2v3h3v2Z"/>`),
		g.Group(children),
	)
}