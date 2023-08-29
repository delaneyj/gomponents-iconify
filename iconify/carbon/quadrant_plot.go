package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuadrantPlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 15H17V2h-2v13H2v2h13v13h2V17h13v-2z"/><path fill="currentColor" d="M5 30a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zM8 8a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1zm14 9a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3zm0-4a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1z"/><circle cx="11" cy="11" r="2" fill="currentColor"/><circle cx="11" cy="21" r="2" fill="currentColor"/><circle cx="21" cy="21" r="2" fill="currentColor"/><circle cx="22" cy="28" r="2" fill="currentColor"/><circle cx="28" cy="24" r="2" fill="currentColor"/><circle cx="4" cy="11" r="2" fill="currentColor"/><circle cx="28" cy="4" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}