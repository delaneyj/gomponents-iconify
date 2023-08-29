package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bmp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 16V8h2a2 2 0 1 1 0 4h-2M6 14a2 2 0 0 1-2 2H2V8h2a2 2 0 1 1 0 4H2h2a2 2 0 0 1 2 2zm3 2V8l3 6l3-6v8"/>`),
		g.Group(children),
	)
}