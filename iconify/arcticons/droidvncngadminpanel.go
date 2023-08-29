package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droidvncngadminpanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.363 42.353V31.745H24.01a7.744 7.744 0 1 0 0-15.489H5.363V5.647H24.01a18.353 18.353 0 0 1 0 36.707Z"/>`),
		g.Group(children),
	)
}