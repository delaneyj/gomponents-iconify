package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mykyivstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.5 19.397l12.527 4.07M24 4.503v13.173m7.973 5.792l12.527-4.07M28.927 32.84l7.742 10.657M19.073 32.84l-7.742 10.657"/>`),
		g.Group(children),
	)
}