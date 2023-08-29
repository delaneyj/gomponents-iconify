package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grammarly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.46 33.152a21.5 21.5 0 1 1 .004-18.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m44.759 39.447l-1.299-6.295l-6.294 1.299"/>`),
		g.Group(children),
	)
}