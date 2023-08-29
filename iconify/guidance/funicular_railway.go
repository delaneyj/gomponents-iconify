package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunicularRailway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.478 18.149a1.5 1.5 0 0 1-2.954.52m11.999-2.25a1.5 1.5 0 0 0 2.954-.52M8 11.758V4.636m8 5.648V3.182m6.97 6.23c.019-.477.03-.98.03-1.503C23 4.41 22.5 2 22.5 2l-21 3.818S1 8.41 1 11.91c0 .523.011 1.022.03 1.492m21.94-3.99C22.862 12.127 22.5 14 22.5 14l-21 3.818s-.362-1.743-.47-4.417m21.94-3.99c-10.656.973-21.302 3.818-21.94 3.99M23 19L1 23"/>`),
		g.Group(children),
	)
}