package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v12a2 2 0 0 0 2 2h4l4 4l4-4h4c1.11 0 2-.89 2-2V4a2 2 0 0 0-2-2m-4.46 10.12l-1.42 1.42L12 11.41l-2.12 2.13l-1.41-1.42L10.59 10L8.46 7.88l1.42-1.41L12 8.59l2.12-2.12l1.42 1.41L13.41 10l2.13 2.12Z"/>`),
		g.Group(children),
	)
}