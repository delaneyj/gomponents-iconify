package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lumicall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.007 10.071L16.293 4.5l15.6 26l4.643-2.786l4.457 7.429L27.064 43.5Z"/>`),
		g.Group(children),
	)
}