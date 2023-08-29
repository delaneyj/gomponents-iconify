package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Siddharthabanksmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 3.495l20.499 20.499l-20.5 20.499l-20.498-20.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.83 32.5c2.37 0 4.53-4.78 5.17-8.5s2.8-8.5 5.17-8.5"/>`),
		g.Group(children),
	)
}