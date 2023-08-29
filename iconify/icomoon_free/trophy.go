package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 3V1H3v2H0v2a3 3 0 0 0 3.9 2.862A5 5 0 0 0 7 9.899v3.1H6a2 2 0 0 0-2 2h8a2 2 0 0 0-2-2H9v-3.1a5.003 5.003 0 0 0 3.1-2.037A3 3 0 0 0 16 5V3h-3zM3 6.813A1.815 1.815 0 0 1 1.187 5V4H3v1c0 .628.116 1.229.327 1.782c-.106.019-.216.03-.327.03zM14.813 5a1.815 1.815 0 0 1-2.14 1.783A4.994 4.994 0 0 0 13 5.001v-1h1.813v1z"/>`),
		g.Group(children),
	)
}