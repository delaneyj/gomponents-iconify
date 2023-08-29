package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightBallpool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.8 20.65a3.38 3.38 0 0 0 3-3.75h0a3.39 3.39 0 0 0-3.75-3l-2.19.25a3.4 3.4 0 0 0-3 3.76h0a3.4 3.4 0 0 0 3.76 3l2.18-.26l-2.18.26a3.39 3.39 0 0 0-3 3.75h0a3.4 3.4 0 0 0 3.76 3l2.18-.26a3.39 3.39 0 0 0 3-3.75h0a3.4 3.4 0 0 0-3.76-3"/><circle cx="23.1" cy="21.53" r="12.09" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}