package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.143 14.255l-3.52 13.139a1.833 1.833 0 0 1-1.771 1.358h-9.018a1.833 1.833 0 0 1-1.77-2.307l1.513-5.648a1.833 1.833 0 0 1 1.77-1.358h8.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.393 19.439l-3.375 12.595a2.308 2.308 0 0 1-2.229 1.71h-8.056m16.771-4.991L37 19.439"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}