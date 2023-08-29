package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.286 11.3a1 1 0 0 0 1.41 1.419"/><path d="M8.44 8.49a5 5 0 0 0 7.098 7.044m1.377-2.611a5 5 0 0 0-5.846-5.836"/><path d="M5.649 5.623a9 9 0 1 0 12.698 12.758m1.683-2.313A9 9 0 0 0 7.954 3.958M3 3l18 18"/></g>`),
		g.Group(children),
	)
}