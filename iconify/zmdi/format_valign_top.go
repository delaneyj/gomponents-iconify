package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatValignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m85 171l86-86l85 86h-64v213h-43V171H85zM0 0h341v43H0V0z"/>`),
		g.Group(children),
	)
}