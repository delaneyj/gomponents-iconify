package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcbSafekey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.316 16.298c1.31 0 2.37 1.078 2.37 2.409s-1.06 2.408-2.37 2.408h-3.912V11.48h3.912c1.31 0 2.37 1.079 2.37 2.409s-1.06 2.409-2.37 2.409Zm0 0h-3.912m-2.852 1.586v.04c0 1.762-1.406 3.191-3.141 3.191s-3.141-1.429-3.141-3.191v-3.252c0-1.763 1.406-3.192 3.141-3.192h0c1.735 0 3.141 1.429 3.141 3.192v.04m-10.186 3.18H16.18m-1.044 3.196l3.141-9.608l3.14 9.637M5.5 19.188a6.167 6.167 0 1 1 0 12.334m6.167-6.167H42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}