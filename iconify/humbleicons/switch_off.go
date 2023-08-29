package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M12 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path d="M3 12a6 6 0 0 0 6 6h6a6 6 0 0 0 0-12H9a6 6 0 0 0-6 6Z"/></g>`),
		g.Group(children),
	)
}