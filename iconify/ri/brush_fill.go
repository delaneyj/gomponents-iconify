package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.289 6.213l4.939-3.842a1 1 0 0 1 1.32.083l2.995 2.994a1 1 0 0 1 .082 1.32l-3.84 4.939a7.505 7.505 0 0 1-7.283 9.292C8 20.998 3.5 19.496 1 17.996c3.98-3 3.047-4.81 3.5-6.5c1.058-3.95 4.842-6.258 8.789-5.284ZM16.7 8.09c.066.064.13.129.194.194L18.03 9.42l2.475-3.182l-1.746-1.746l-3.182 2.475L16.7 8.09Z"/>`),
		g.Group(children),
	)
}