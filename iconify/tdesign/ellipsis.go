package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ellipsis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10.5h3v3H3v-3Zm7.5 0h3v3h-3v-3Zm7.5 0h3v3h-3v-3Z"/>`),
		g.Group(children),
	)
}