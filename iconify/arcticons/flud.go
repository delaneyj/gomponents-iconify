package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.66 4.5c0 4.09-13 15.8-13 25.58s7.67 13.42 13 13.42c4.77 0 13.64-2 13.64-12.62c0-11.6-13.64-18.53-13.64-26.38Z"/>`),
		g.Group(children),
	)
}