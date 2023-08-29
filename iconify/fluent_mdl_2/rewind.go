package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1674L90 1024l806-645v1295zm-128-267V645l-474 379l474 383zm218-383l806-645v1295l-806-650zm678 383V645l-474 379l474 383z"/>`),
		g.Group(children),
	)
}