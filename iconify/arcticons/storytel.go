package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storytel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.67 43.5c4.47-5.72 32.81-4.74 32.81-22.5c0-12.28-9.7-16.5-16.48-16.5S6.52 10.42 6.52 26.42c0 12.43 2.15 17.08 2.15 17.08Z"/>`),
		g.Group(children),
	)
}