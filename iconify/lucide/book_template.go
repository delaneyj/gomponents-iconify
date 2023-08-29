package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 22h-2m2-7v2h-2M4 19.5V15m16-7v3m-2-9h2v2M4 11V9m8-7h2m-2 20h2m-2-5h2m-6 5H6.5a2.5 2.5 0 0 1 0-5H8M4 5v-.5A2.5 2.5 0 0 1 6.5 2H8"/>`),
		g.Group(children),
	)
}