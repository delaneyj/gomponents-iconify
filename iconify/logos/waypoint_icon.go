package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaypointIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#14C6CB" d="m256 0l-31.961 55.366L192.078 0H256ZM85.344 36.952H64.016l53.32 92.381l-21.328 36.952L0 0h106.672l53.351 92.38l10.664-18.475L128 0h42.687l21.328 36.952l21.329 36.953l-53.32 92.38l-74.68-129.333Z"/>`),
		g.Group(children),
	)
}