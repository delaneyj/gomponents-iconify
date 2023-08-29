package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20h4V4h-4v16m-6 0h4v-8H4v8M16 9v11h4V9h-4Z"/>`),
		g.Group(children),
	)
}