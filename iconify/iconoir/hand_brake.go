package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandBrake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 16v-4m0-3V8"/><circle cx="12" cy="12" r="8"/><path stroke-linecap="round" stroke-linejoin="round" d="M3.953 4.5A10.961 10.961 0 0 0 1 12c0 2.899 1.121 5.535 2.953 7.5m16.094-15A10.962 10.962 0 0 1 23 12c0 2.899-1.121 5.535-2.953 7.5"/></g>`),
		g.Group(children),
	)
}