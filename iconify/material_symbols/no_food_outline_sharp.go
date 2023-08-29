package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFoodOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425Zm1.15-4.55L19.8 16.925L20.8 7h-9.55L11 5h5V1h2v4h5l-1.375 13.75Zm-6-5.975ZM1 19v-2h15v2H1Zm0 4v-2h15v2H1ZM9.05 9.025v2q-.125 0-.275-.013T8.5 11q-1.475 0-2.788.5T3.676 13h9.35l2 2H1q0-3.025 2.337-4.513T8.5 9q.125 0 .275.013t.275.012ZM8.5 13Z"/>`),
		g.Group(children),
	)
}