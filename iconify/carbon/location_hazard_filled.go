package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationHazardFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M16 22a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 22Zm-1.125-6h2.25V7h-2.25Z"/><path fill="currentColor" d="M16 2A11.013 11.013 0 0 0 5 13a10.889 10.889 0 0 0 2.216 6.6s.3.395.349.452L16 30l8.439-9.953c.044-.053.345-.447.345-.447l.001-.003A10.885 10.885 0 0 0 27 13A11.013 11.013 0 0 0 16 2Zm-1.125 5h2.25v9h-2.25ZM16 22a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 22Z"/>`),
		g.Group(children),
	)
}