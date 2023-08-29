package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m29 35l-8-3.667M29 13l-8 3.667m0 0L17 18.5L5 24l12 5.5l4 1.833m0-14.666v14.666M37 6v36l6-6"/>`),
		g.Group(children),
	)
}