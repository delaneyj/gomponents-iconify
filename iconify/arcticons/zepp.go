package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zepp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.874 5.5H8.131c1.264 7.02 2.363 15.061 21.027 18.679C19.278 26.976 7.84 29.837 8.13 42.5h31.743"/>`),
		g.Group(children),
	)
}