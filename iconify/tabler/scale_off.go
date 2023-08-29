package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20h10M9.452 5.425L12 5l6 1m-6-3v5m0 4v8m-3-8L6 6l-3 6a3 3 0 0 0 6 0m9.873 2.871A3 3 0 0 0 21 12l-3-6l-2.677 5.355M3 3l18 18"/>`),
		g.Group(children),
	)
}