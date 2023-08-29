package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motomarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.33 4.5H15.67a3.16 3.16 0 0 0-3.14 3.14v32.72a3.16 3.16 0 0 0 3.14 3.14h16.66a3.16 3.16 0 0 0 3.14-3.14V7.64a3.16 3.16 0 0 0-3.14-3.14ZM24 20.22a6 6 0 1 1 6-6a6 6 0 0 1-6 6Z"/><circle cx="24" cy="14.27" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}