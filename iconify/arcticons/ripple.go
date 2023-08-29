package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ripple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5A14.27 14.27 0 0 1 9.75 29.32C9.75 19.89 24 4.5 24 4.5s14.25 15.39 14.25 24.82A14.27 14.27 0 0 1 24 43.5Z"/>`),
		g.Group(children),
	)
}