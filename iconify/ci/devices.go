package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.5v-.3c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C7.52 4 8.08 4 9.2 4h8.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v6.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.874.875C19.48 17 18.92 17 17.803 17H10.5M3 16.8v-5.6c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 8 5.08 8 6.2 8h.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875C8.48 20 7.921 20 6.803 20h-.606c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 18.48 3 17.92 3 16.8Z"/>`),
		g.Group(children),
	)
}