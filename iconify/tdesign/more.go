package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func More(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 3h3v3h-3V3Zm0 7.5h3v3h-3v-3Zm0 7.5h3v3h-3v-3Z"/>`),
		g.Group(children),
	)
}