package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 1v6h2V3h4V1zm2 12H1v6h6v-2H3zm14 4h-4v2h6v-6h-2zm0-16h-4v2h4v4h2V1z"/>`),
		g.Group(children),
	)
}