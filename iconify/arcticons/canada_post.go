package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CanadaPost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.01 24.935L4.5 10.792l39 13.325L30.27 36.63l-25.77.579l11.51-12.273Zm20.265-.461l-20.266.461m16.442 3.641l-20.267.461m16.443 3.641l-20.269.461"/>`),
		g.Group(children),
	)
}