package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9h4l5-5v16l-5-5H3V9m11 2h3V8h2v3h3v2h-3v3h-2v-3h-3v-2Z"/>`),
		g.Group(children),
	)
}