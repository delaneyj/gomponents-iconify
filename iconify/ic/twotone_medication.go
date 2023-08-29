package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneMedication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19h10V8H7v11zm1-7h2.5V9.5h3V12H16v3h-2.5v2.5h-3V15H8v-3z" opacity=".3"/><path fill="currentColor" d="M6 3h12v2H6zm11 3H7c-1.1 0-2 .9-2 2v11c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2zm0 13H7V8h10v11z"/><path fill="currentColor" d="M10.5 17.5h3V15H16v-3h-2.5V9.5h-3V12H8v3h2.5z"/>`),
		g.Group(children),
	)
}