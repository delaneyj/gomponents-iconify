package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadicalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 72v24a6 6 0 0 1-12 0V78H124.16L77.62 202.11a6 6 0 0 1-11.24 0l-48-128a6 6 0 0 1 11.24-4.22L72 182.91l42.38-113A6 6 0 0 1 120 66h112a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}