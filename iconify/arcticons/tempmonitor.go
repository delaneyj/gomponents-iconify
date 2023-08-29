package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tempmonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.6 28V9.3a4.74 4.74 0 0 0-4.13-4.79a4.61 4.61 0 0 0-5.07 4.57V28a1 1 0 0 1-.4.8a8.19 8.19 0 1 0 10 0a1 1 0 0 1-.4-.8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.4 26.18c2.17-1.16 4.37-.64 6-.42a7.92 7.92 0 0 0 3.2-.32"/>`),
		g.Group(children),
	)
}