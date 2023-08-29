package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.57 54.42a104 104 0 1 0 0 147.15a104.17 104.17 0 0 0 0-147.15Zm-75.91 79.24l-32 32a8 8 0 1 1-11.32-11.32L108.68 128l-26.36-26.36a8 8 0 1 1 11.32-11.32l32 32a8 8 0 0 1 .02 11.34Zm56 0l-32 32a8 8 0 0 1-11.32-11.32L164.71 128l-26.36-26.36a8 8 0 1 1 11.32-11.32l32 32a8 8 0 0 1 .01 11.34Z"/>`),
		g.Group(children),
	)
}