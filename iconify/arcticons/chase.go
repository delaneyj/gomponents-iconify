package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 30.82l10.27 11.07a1.89 1.89 0 0 0 1.39.61h13.66V30.82ZM17.18 5.5L6.11 15.77a1.89 1.89 0 0 0-.61 1.39v13.66h11.68ZM42.5 17.18L32.23 6.11a1.89 1.89 0 0 0-1.39-.61H17.18v11.68ZM30.82 42.5l11.07-10.27a1.89 1.89 0 0 0 .61-1.39V17.18H30.82Z"/>`),
		g.Group(children),
	)
}