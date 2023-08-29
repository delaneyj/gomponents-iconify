package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatValignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m256 213l-85 86l-86-86h64V0h43v213h64zM0 341h341v43H0v-43z"/>`),
		g.Group(children),
	)
}