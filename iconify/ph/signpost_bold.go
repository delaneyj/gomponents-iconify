package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m248.92 104l-36-40a12 12 0 0 0-8.92-4h-64V32a12 12 0 0 0-24 0v28H40a20 20 0 0 0-20 20v64a20 20 0 0 0 20 20h76v60a12 12 0 0 0 24 0v-60h64a12 12 0 0 0 8.92-4l36-40a12 12 0 0 0 0-16Zm-50.26 36H44V84h154.66l25.2 28Z"/>`),
		g.Group(children),
	)
}