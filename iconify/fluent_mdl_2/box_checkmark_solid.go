package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxCheckmarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 0v1920H0V0h1920zm-358 621l-135-135l-659 658l-275-274l-135 135l410 411l794-795z"/>`),
		g.Group(children),
	)
}