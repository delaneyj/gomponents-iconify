package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v10h-2V4H4v9.586l5-5L14.414 14L13 15.414l-4-4l-5 5V20h8v2H2V2Zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0ZM20 14v4h4v2h-4v4h-2v-4h-4v-2h4v-4h2Z"/>`),
		g.Group(children),
	)
}