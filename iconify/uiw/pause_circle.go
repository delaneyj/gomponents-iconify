package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0ZM8.056 5.455a.682.682 0 0 0-.682.681v7.272a.682.682 0 0 0 1.364 0V6.136a.682.682 0 0 0-.682-.681Zm4.546 0a.682.682 0 0 0-.682.681v7.272a.682.682 0 0 0 1.363 0V6.136a.682.682 0 0 0-.681-.681Z"/>`),
		g.Group(children),
	)
}