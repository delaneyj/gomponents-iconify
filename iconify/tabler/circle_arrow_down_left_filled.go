package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowDownLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 3.34a10 10 0 1 1-14.995 8.984L2 12l.005-.324A10 10 0 0 1 17 3.34zM9 8a1 1 0 0 0-1 1v6l.007.117l.029.149l.035.105l.054.113l.071.111c.03.04.061.077.097.112l.09.08l.096.067l.098.052l.11.044l.112.03l.126.017L15 16l.117-.007A1 1 0 0 0 16 15l-.007-.117A1 1 0 0 0 15 14h-3.586l4.293-4.293l.083-.094a1 1 0 0 0-1.497-1.32L10 12.584V9l-.007-.117A1 1 0 0 0 9 8z"/></g>`),
		g.Group(children),
	)
}