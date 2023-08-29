package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAscending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m43.314 130.51l51.882-51.883v417.569h32V78.627l51.883 51.883l22.627-22.627l-90.51-90.511l-90.51 90.511l22.628 22.627z"/><path fill="currentColor" d="M184 160h120v32H184zm0 72h184v32H184zm0 72h248v32H184zm0 72h312v32H184z"/>`),
		g.Group(children),
	)
}