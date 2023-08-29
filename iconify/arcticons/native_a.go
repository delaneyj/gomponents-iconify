package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NativeA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.434 7.298S27.715 40.34 15.692 40.702C1.702 40.54 1.823 7.419 17.862 7.298C31.852 7.503 28.39 40.22 37.4 40.1c1.604.024 3.063-1.077 5.101-5.029"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.77 28.667v-10.01l8.321 10.01v-10.01"/>`),
		g.Group(children),
	)
}