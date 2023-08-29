package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M21 8V6H3v14H2v2h6v-7c0-.163.046-4 4-4c3.821 0 3.993 3.602 4 4v7h6v-2h-1V8zM2 2h20v2H2z" fill="currentColor"/>`),
		g.Group(children),
	)
}