package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebpackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 10.5v-6m0 6l6 3.5m-6-3.5L5 9M1.5 4.5l6-3.5m-6 3.5l6 3m0-6.5l6 3.5M7.5 1v3.5m6 0v6m0-6l-6 3m6 3l-6 3.5m6-3.5L10 9m-2.5 5V7.5m0-3L5 6v3m2.5-4.5L10 6v3M5 9l2.5 1.5L10 9"/>`),
		g.Group(children),
	)
}