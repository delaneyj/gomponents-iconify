package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchbarEx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="18.451" cy="18.451" r="12.951" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.61 27.61l1.507 1.507m0 0v3.771l9.612 9.612l3.771-3.771l-9.612-9.612h-3.771z"/>`),
		g.Group(children),
	)
}