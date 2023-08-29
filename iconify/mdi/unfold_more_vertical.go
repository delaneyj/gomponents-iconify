package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.17 12L15 8.83l1.41-1.42L21 12l-4.59 4.58L15 15.17L18.17 12M5.83 12L9 15.17l-1.41 1.42L3 12l4.59-4.58L9 8.83L5.83 12Z"/>`),
		g.Group(children),
	)
}