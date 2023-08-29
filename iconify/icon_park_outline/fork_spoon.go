package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkSpoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 4v40M8 5v10c0 5 6 5 6 5s6 0 6-5V5m14 15v24m6-32c0 4.418-2.686 8-6 8s-6-3.582-6-8s2.686-8 6-8s6 3.582 6 8Z"/>`),
		g.Group(children),
	)
}