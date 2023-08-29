package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFooterTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.746 1.996a2.25 2.25 0 0 1 2.245 2.096l.005.154v15.498A2.25 2.25 0 0 1 17.9 21.99l-.154.005h-11.5a2.25 2.25 0 0 1-2.245-2.096l-.005-.154V4.246a2.25 2.25 0 0 1 2.096-2.245l.154-.005h11.5ZM8.501 16a1.5 1.5 0 0 0 0 3h7a1.5 1.5 0 0 0 0-3h-7Z"/>`),
		g.Group(children),
	)
}