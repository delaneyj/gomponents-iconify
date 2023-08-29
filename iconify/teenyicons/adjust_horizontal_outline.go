package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustHorizontalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5 7.5H0m15 5h-2.5m2.5-10H8.5m-2 0H0m4.5 5H15m-4.5 5H0m10.5-2v4h2v-4h-2Zm-8-5v4h2v-4h-2Zm4-5v4h2v-4h-2Z"/>`),
		g.Group(children),
	)
}