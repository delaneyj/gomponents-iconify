package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitRepository(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5.5A4.5 4.5 0 0 1 7.5 1H21v22H7.5A4.5 4.5 0 0 1 3 18.5v-13Zm2 9.258A4.479 4.479 0 0 1 7.5 14H19V3H7.5A2.5 2.5 0 0 0 5 5.5v9.258ZM8 16h-.5a2.5 2.5 0 0 0 0 5H19v-5h-5v4.618l-3-1.5l-3 1.5V16Zm4 0h-2v1.382l1-.5l1 .5V16ZM8 5h2.004v2.004H8V5Zm0 3h2.004v2.004H8V8Z"/>`),
		g.Group(children),
	)
}