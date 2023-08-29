package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/></g>`),
		g.Group(children),
	)
}