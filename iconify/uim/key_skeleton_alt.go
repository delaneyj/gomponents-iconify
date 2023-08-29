package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySkeletonAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 12a5 5 0 1 0 5 5a5 5 0 0 0-5-5zm-1.415 7.412a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M21.708 6.534L20.293 5.12l1.413-1.413a1 1 0 1 0-1.414-1.414L9.753 12.831a5.018 5.018 0 0 1 1.415 1.414l4.883-4.883l1.414 1.414a1 1 0 0 0 1.414-1.414l-1.414-1.414l1.414-1.414l1.415 1.414a1 1 0 0 0 1.414-1.414z"/>`),
		g.Group(children),
	)
}