package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v4h4V4H4Zm6 0v4h4V4h-4Zm6 0v4h4V4h-4Zm4 6h-4v4h4v-4Zm0 6h-4v4h4v-4Zm-6 4v-4h-4v4h4Zm-6 0v-4H4v4h4Zm-4-6h4v-4H4v4Zm6-4v4h4v-4h-4Z"/>`),
		g.Group(children),
	)
}