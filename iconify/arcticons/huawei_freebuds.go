package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiFreebuds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.904 41.197l-5.803-10.682a4.463 4.463 0 0 1-.27-4.417c1.858-4.214-.195-8.34-3.99-9.33c-4.27-1.116-7.68.68-8.991 2.569c-1.408 2.026-3.316 4.549 1.014 8.587a7.965 7.965 0 0 0 4.89 1.803c1.677.135 3.172.744 3.832 2.456l5.5 11.303"/>`),
		g.Group(children),
	)
}