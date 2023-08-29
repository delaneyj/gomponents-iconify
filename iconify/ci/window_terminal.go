package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowTerminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16m-4 10h-3m-5-4l2 2l-2 2m-5 .8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 18.48 3 17.92 3 16.8Z"/>`),
		g.Group(children),
	)
}