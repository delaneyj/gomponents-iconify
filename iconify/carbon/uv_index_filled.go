package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UvIndexFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 30l-3.463-5.822L6 26l1.822-6.537L2 16l5.822-3.463L6 6l6.537 1.822L16 2l3.463 5.822L26 6l-1.822 6.537L30 16l-5.822 3.463L26 26l-6.537-1.822Z"/>`),
		g.Group(children),
	)
}