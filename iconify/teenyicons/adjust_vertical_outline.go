package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustVerticalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 12.5V15m5-15v2.5M2.5 0v6.5m0 2V15m5-4.5V0m5 4.5V15m-2-10.5h4v-2h-4v2Zm-5 8h4v-2h-4v2Zm-5-4h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}