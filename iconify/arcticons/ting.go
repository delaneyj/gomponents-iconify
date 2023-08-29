package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.758 38.935h.772a2.146 2.146 0 0 0 2.146-2.145V8.954a2.146 2.146 0 0 0-2.146-2.146H7.646A2.146 2.146 0 0 0 5.5 8.954V36.79a2.146 2.146 0 0 0 2.146 2.146h4.442M42.5 6.809H26.958a2.146 2.146 0 0 0-2.146 2.145v32.237m0-22.56H42.5m-3.068 22.56v-22.56"/>`),
		g.Group(children),
	)
}