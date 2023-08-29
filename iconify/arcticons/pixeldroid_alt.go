package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PixeldroidAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.219 5.713H11.86A3.167 3.167 0 0 0 8.694 8.88v32.404a1 1 0 0 0 1.742.671l9.274-10.25a2 2 0 0 1 1.483-.658h5.446a12.667 12.667 0 0 0 12.665-12.904a12.943 12.943 0 0 0-13.085-12.43Z"/>`),
		g.Group(children),
	)
}