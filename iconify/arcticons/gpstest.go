package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpstest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 41.06l9.1-12L21.26 37l9.1-11.8L42.5 41.06Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.37 6.94A11.61 11.61 0 0 1 5.76 18.55V6.94Z"/><path d="M11.57 6.94a5.81 5.81 0 0 1-5.81 5.8v-5.8Z"/></g>`),
		g.Group(children),
	)
}