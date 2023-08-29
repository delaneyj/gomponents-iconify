package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerSandPaused(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 6v12h-6l-4-4l-4 4H2V6h6l4 4l4-4M7.5 16l4-4l-4-4H4v8m8.5-4l4 4H20V8h-3.5m1.5 4v2h-.8l-2-2m-6.4 0l-2 2H6v-2Z"/>`),
		g.Group(children),
	)
}