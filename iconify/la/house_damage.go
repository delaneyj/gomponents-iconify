package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseDamage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 2.59l-.72.69l-13 13l1.44 1.44L5 16.44V28h8.83L16 22h-3.62l3.63-6.21L16 20h3.55l-1.61 8H27V16.44l1.28 1.28l1.44-1.44l-13-13l-.72-.69zm0 2.85l9 9V26h-4.62L22 18h-4v-6h-2.1L10 22.11V24h3.15l-.72 2H7V14.44l9-9z"/>`),
		g.Group(children),
	)
}