package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerSandEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2v6l4 4l-4 4v6h12v-6l-4-4l4-4V2H6m10 14.5V20H8v-3.5l4-4l4 4m-4-5l-4-4V4h8v3.5l-4 4Z"/>`),
		g.Group(children),
	)
}