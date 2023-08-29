package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.849 9.245c-.417-.518-2.513-2.1-1.658-4.219c0 0-1.908.725-1.149 2.953c-.35-.496-3.318-1.616-3.04-7.85c0 0-4.912 2.954-4.912 7.059c0 0 0 3.256 1.974 4.799c0 0-2.616-.43-3.317-4.557c-.176.494-4.756 5.288 1.74 8.446l3.489.083h3.351s7.372-2.493 3.522-6.714z"/>`),
		g.Group(children),
	)
}