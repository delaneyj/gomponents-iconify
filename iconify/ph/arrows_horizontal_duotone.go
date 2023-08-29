package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsHorizontalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m24 128l32-32v64Zm176-32v64l32-32Z" opacity=".2"/><path d="m237.66 122.34l-32-32A8 8 0 0 0 192 96v24H64V96a8 8 0 0 0-13.66-5.66l-32 32a8 8 0 0 0 0 11.32l32 32A8 8 0 0 0 64 160v-24h128v24a8 8 0 0 0 13.66 5.66l32-32a8 8 0 0 0 0-11.32ZM48 140.69L35.31 128L48 115.31Zm160 0v-25.38L220.69 128Z"/></g>`),
		g.Group(children),
	)
}