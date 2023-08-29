package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnaTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 3v1c-.01 3.352-1.68 6.023-5.008 8.014c-3.328 1.99 3.336-2 .008-.014c-3.328 1.99-5 4.662-5.008 8.014v1"/><path d="M17 21.014v-1c-.01-3.352-1.68-6.023-5.008-8.014c-3.328-1.99 3.336 2 .008.014C8.672 10.023 7 7.352 6.992 4V3M7 4h10M7 20h10M8 8h8m-8 8h8"/></g>`),
		g.Group(children),
	)
}