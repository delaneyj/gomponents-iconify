package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h6v2H5v4H3V3Zm0 18h6v-2H5v-4H3v6Zm12 0h6v-6h-2v4h-4v2Zm6-18h-6v2h4v4h2V3Z"/>`),
		g.Group(children),
	)
}