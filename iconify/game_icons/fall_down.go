package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M257.75 16.03A60 60 0 0 0 196 76a60 60 0 0 0 120 0a60 60 0 0 0-58.25-59.97zM250.72 166c-24.72.11-24.72 1.875-24.72 30v210h-60l90 90l90-90h-60V196c0-30 0-30-30-30c-1.875 0-3.633-.007-5.28 0z"/>`),
		g.Group(children),
	)
}