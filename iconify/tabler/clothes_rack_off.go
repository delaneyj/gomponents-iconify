package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesRackOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2 2v1m0 4v9m-3 0h6M7.757 9.243a6 6 0 0 0 3.129 1.653m3.578-.424a6 6 0 0 0 1.779-1.229M3 3l18 18"/>`),
		g.Group(children),
	)
}