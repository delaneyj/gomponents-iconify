package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acorns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.808 22.467c-1.592 7.905 1.1 17.893 11.405 19.233c.136.018 1.012 1.078 1.8 1.8l1.8-1.8c12.218-2.218 11.808-14.736 11.226-19.12"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.31 22.071C6.67 17.188 12.117 9.386 22.405 9.05a5.948 5.948 0 0 0 .118-4.55a8.736 8.736 0 0 1 3.578 4.653c11.379.53 14.683 8.85 12.718 12.926c-3.433 1.573-22.802 2.156-29.51-.008Z"/>`),
		g.Group(children),
	)
}