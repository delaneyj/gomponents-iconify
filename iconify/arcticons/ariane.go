package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ariane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="24" cy="23.99" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.11" ry="21.48" transform="rotate(-45 24 23.999)"/>`),
		g.Group(children),
	)
}