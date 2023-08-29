package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.83 8.77l-2.77 2.84H6.29A1.79 1.79 0 0 0 4.5 13.4v23.22a1.8 1.8 0 0 0 1.79 1.8h35.42a1.8 1.8 0 0 0 1.79-1.8V13.4a1.79 1.79 0 0 0-1.79-1.79H30.94l-2.77-2.84Zm18.93 5.74a1.84 1.84 0 1 1 0 3.68a1.84 1.84 0 0 1 0-3.68ZM24 17.71a8.51 8.51 0 1 1-8.51 8.51A8.51 8.51 0 0 1 24 17.71Z"/>`),
		g.Group(children),
	)
}