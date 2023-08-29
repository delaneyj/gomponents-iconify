package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.146 19.996a8.146 8.146 0 0 0-8.584-8.134c-4.405.23-7.708 4.184-7.708 8.595v7.547A8.146 8.146 0 0 0 24 36.15h0a8.146 8.146 0 0 0 8.146-8.146v-8.008m-1.934-5.265L17.779 33.259"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}