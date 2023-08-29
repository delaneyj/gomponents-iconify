package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15h1M9 15h6m-9 0H5m0-3h14M5 9h14M2 14.8V9.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C3.52 6 4.08 6 5.2 6h13.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H5.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C2 16.48 2 15.92 2 14.8Z"/>`),
		g.Group(children),
	)
}