package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Task(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v20H3V3h4V1Zm0 4H5v16h14V5h-2v2H7V5Zm8-2H9v2h6V3Zm-6 8h6v2H9v-2Zm0 4h6v2H9v-2Z"/>`),
		g.Group(children),
	)
}