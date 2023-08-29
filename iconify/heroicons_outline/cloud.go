package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 0 0 4 4h9a5 5 0 1 0-.1-9.999a5.002 5.002 0 1 0-9.78 2.096A4.001 4.001 0 0 0 3 15Z"/>`),
		g.Group(children),
	)
}