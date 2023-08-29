package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSubUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 13l5-5m0 0l-5-5m5 5h-7.8c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C7 9.52 7 10.08 7 11.2V21"/>`),
		g.Group(children),
	)
}