package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19h14m1.986-2.014a2 2 0 0 0-.146-.736L13.74 4a2 2 0 0 0-3.5 0l-.825 1.424m-1.467 2.53L3.14 16.25A2 2 0 0 0 4.89 19M3 3l18 18"/>`),
		g.Group(children),
	)
}