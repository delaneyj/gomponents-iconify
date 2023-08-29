package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9H3V7h4V3h2v6Zm0 6H3v2h4v4h2v-6Zm12 0h-6v6h2v-4h4v-2Zm-6-6h6V7h-4V3h-2v6Z"/>`),
		g.Group(children),
	)
}