package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterPercentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .269L5.636 6.633a9 9 0 1 0 12.728 0L12 .269Zm4.243 9.902l-7.071 7.072l-1.415-1.415l7.072-7.07l1.414 1.413ZM8.11 9.111a1.5 1.5 0 1 1 2.121 2.121A1.5 1.5 0 0 1 8.11 9.111Zm7.778 7.778a1.5 1.5 0 1 1-2.121-2.121a1.5 1.5 0 0 1 2.121 2.12Z"/>`),
		g.Group(children),
	)
}