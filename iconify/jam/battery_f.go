package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 0h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path d="M18.5 3h.5a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-.5A1.5 1.5 0 0 1 17 7.5v-3A1.5 1.5 0 0 1 18.5 3z"/></g>`),
		g.Group(children),
	)
}