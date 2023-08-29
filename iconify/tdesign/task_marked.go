package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v8h-2V5h-2v2H7V5H5v16h7.5v2H3V3h4V1Zm2 4h6V3H9v2Zm5.75 7.996h8.5v10.295l-4.247-2.617l-4.253 2.615V12.996Zm2 2v4.715l2.254-1.385l2.246 1.383v-4.713h-4.5Z"/>`),
		g.Group(children),
	)
}