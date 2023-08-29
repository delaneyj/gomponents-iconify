package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cafeteria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5 20.5h12.5v-7a3 3 0 1 1 3 3H20m-20 7h24m-9-13H2.5v.25l.063.1A19.008 19.008 0 0 1 5.5 21m0-21v1A6.5 6.5 0 0 0 12 7.5m2.5.5V4.5H12A3.5 3.5 0 0 1 8.5 1V0"/>`),
		g.Group(children),
	)
}