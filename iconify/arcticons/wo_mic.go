package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WoMic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5h0a8.503 8.503 0 0 1 8.503 8.503v10.512A8.503 8.503 0 0 1 24 32.017h0a8.503 8.503 0 0 1-8.503-8.502V13.003A8.503 8.503 0 0 1 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.855 26.045a13.387 13.387 0 0 0 26.29 0M24 36.898V43.5m-13.145 0h26.29"/>`),
		g.Group(children),
	)
}