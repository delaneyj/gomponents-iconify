package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.25 17h-.053c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 15.48 3 14.92 3 13.8V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.874.875C19.48 17 18.92 17 17.803 17h-.05M16 20H8l4-5l4 5Z"/>`),
		g.Group(children),
	)
}