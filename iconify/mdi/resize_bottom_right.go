package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22h-2v-2h2v2m0-4h-2v-2h2v2m-4 4h-2v-2h2v2m0-4h-2v-2h2v2m-4 4h-2v-2h2v2m8-8h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}