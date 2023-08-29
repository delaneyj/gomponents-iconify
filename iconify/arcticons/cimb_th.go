package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CimbTh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.604 24L4.5 8.267h26.942L43.5 24H16.604L4.5 39.733h26.942L43.5 24"/>`),
		g.Group(children),
	)
}