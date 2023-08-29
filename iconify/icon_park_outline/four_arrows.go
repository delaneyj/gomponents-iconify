package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 24l-5-5m5-11v16V8Zm0 16l5-5l-5 5Zm0 0l-5 5m5 11V24v16Zm0-16l5 5l-5-5Zm-4-12a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-8h32m0 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM28 44a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}