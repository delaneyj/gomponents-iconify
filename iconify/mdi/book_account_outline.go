package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookAccountOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h12m0 2h-5v5l-2.5-2.3L8 9V4H6v16h12m-5-9a2 2 0 1 1-2 2a2 2 0 0 1 2-2m4 8H9v-1c0-1.33 2.67-2 4-2s4 .67 4 2v1"/>`),
		g.Group(children),
	)
}