package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wnba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.248 20l-3.41 8l-.59-8l-3.41 8l-.59-8m8.987 8l1.41-8l3.89 8l1.41-8m5.413 4c1.105 0 1.842.895 1.648 2s-1.248 2-2.353 2h-3.3l1.41-8h3.3c1.105 0 1.843.895 1.648 2c-.195 1.105-1.248 2-2.353 2h0Zm0 0h-3.3m11.283 1.35h-3.544M33.452 28l4.061-8l1.239 8"/>`),
		g.Group(children),
	)
}