package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m237.66 133.66l-32 32A8 8 0 0 1 192 160v-24H64v24a8 8 0 0 1-13.66 5.66l-32-32a8 8 0 0 1 0-11.32l32-32A8 8 0 0 1 64 96v24h128V96a8 8 0 0 1 13.66-5.66l32 32a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}