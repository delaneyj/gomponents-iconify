package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 6H9a6 6 0 1 0 0 12h6a6 6 0 0 0 0-12Z"/><circle cx="15" cy="12" r="3" fill="currentColor"/></g>`),
		g.Group(children),
	)
}