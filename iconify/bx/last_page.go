package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.707 17.707L13.414 12L7.707 6.293L6.293 7.707L10.586 12l-4.293 4.293zM15 6h2v12h-2z"/>`),
		g.Group(children),
	)
}