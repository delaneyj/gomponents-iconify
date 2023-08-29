package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm32-136a6 6 0 0 0-6 6v29.17L99.18 82.91A6 6 0 0 0 90 88v80a6 6 0 0 0 9.18 5.09L154 138.83V168a6 6 0 0 0 12 0V88a6 6 0 0 0-6-6Zm-58 75.17V98.83L148.68 128Z"/>`),
		g.Group(children),
	)
}