package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRotatedOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16.964 16.952l-3.462 3.461c-.782.783-2.222.783-3 0l-6.911-6.91c-.783-.783-.783-2.223 0-3l3.455-3.456m2-2l1.453-1.452c.782-.783 2.222-.783 3 0l6.911 6.91c.783.783.783 2.223 0 3l-1.448 1.45M3 3l18 18"/>`),
		g.Group(children),
	)
}