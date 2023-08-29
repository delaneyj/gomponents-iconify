package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecurityStall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 29h16v15H16zM12 4h24v6H12z"/><path d="M16 10v19m16-19v19"/><path stroke-linejoin="round" d="M4 44h40"/></g>`),
		g.Group(children),
	)
}