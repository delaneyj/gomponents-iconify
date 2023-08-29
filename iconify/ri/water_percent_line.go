package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterPercentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.05 8.047L12 3.097l4.95 4.95a7 7 0 1 1-9.9 0Zm11.314-1.414L12 .269L5.636 6.633a9 9 0 1 0 12.728 0Zm-2.121 3.538l-1.414-1.414l-7.072 7.071l1.415 1.415l7.07-7.072ZM8.11 11.232a1.5 1.5 0 1 0 2.121-2.121a1.5 1.5 0 0 0-2.121 2.121Zm7.778 5.657a1.5 1.5 0 1 1-2.121-2.121a1.5 1.5 0 0 1 2.121 2.12Z"/>`),
		g.Group(children),
	)
}