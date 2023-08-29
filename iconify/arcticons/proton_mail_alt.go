package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonMailAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5h0A15.5 15.5 0 0 1 39.5 20v21.5a2 2 0 0 1-2 2h-27a2 2 0 0 1-2-2V20A15.5 15.5 0 0 1 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 27L22 35.38a4 4 0 0 0 4.26 0L39.5 27M24 10.74h0a9.43 9.43 0 0 1 9.43 9.43v1.73h0h-18.86h0v-1.73A9.43 9.43 0 0 1 24 10.74Z"/>`),
		g.Group(children),
	)
}