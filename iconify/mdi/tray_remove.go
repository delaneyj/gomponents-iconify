package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrayRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5h-2v5H4v-5H2m12.12-6.54l1.42 1.42L13.41 9l2.13 2.12l-1.42 1.42L12 10.41l-2.12 2.13l-1.42-1.42L10.59 9L8.46 6.88l1.42-1.42L12 7.59Z"/>`),
		g.Group(children),
	)
}