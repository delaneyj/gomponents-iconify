package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patreon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="14.508" cy="9.831" r="6.496" fill="currentColor"/><path fill="currentColor" d="M2.996 3.335H6.17v17.33H2.996z"/>`),
		g.Group(children),
	)
}