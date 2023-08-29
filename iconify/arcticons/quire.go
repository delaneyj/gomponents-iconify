package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.338 36.908a17.108 17.108 0 1 1 4.904-23.692"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.302 21.099c-.16 8.45 20.036 7.827 18.81 21.401"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.242 13.216a17.113 17.113 0 0 1-.116 18.971"/>`),
		g.Group(children),
	)
}