package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindmillTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24 24c5.523 0 10-4.477 10-10S29.523 4 24 4v20Zm0 0c0 5.523 4.477 10 10 10s10-4.477 10-10H24Zm0 0c0-5.523-4.477-10-10-10S4 18.477 4 24h20Zm0 0c-5.523 0-10 4.477-10 10s4.477 10 10 10V24Z"/>`),
		g.Group(children),
	)
}