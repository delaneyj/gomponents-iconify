package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonPlayButtonTelevisionButtonsMoviesPlayTvVideoControls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 12.35a1.14 1.14 0 0 0 .63 1a1.24 1.24 0 0 0 1.22 0L12 8a1.11 1.11 0 0 0 0-2L3.35.69a1.24 1.24 0 0 0-1.22 0a1.14 1.14 0 0 0-.63 1Z"/>`),
		g.Group(children),
	)
}