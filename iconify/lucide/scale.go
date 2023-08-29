package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 16l3-8l3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1ZM2 16l3-8l3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Zm5 5h10M12 3v18M3 7h2c2 0 5-1 7-2c2 1 5 2 7 2h2"/>`),
		g.Group(children),
	)
}