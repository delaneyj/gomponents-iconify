package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusNoteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.95 19q-2.5-.125-3.725-1.05T1 15.275q0-1.625 1.338-2.638t3.712-1.212q.975-.075 1.463-.313T8 10.45q0-.65-.738-.975T4.825 9L5 7q2.575.2 3.788 1.038T10 10.45q0 1.325-.963 2.075t-2.837.9q-1.6.125-2.4.588T3 15.275q0 .875.7 1.263T6.05 17l-.1 2Zm7.925-.75L9.75 14.125l9.8-9.8L23.7 8.45l-9.825 9.8Zm-1.15.95l-4.95 1.05l1-5l3.95 3.95Z"/>`),
		g.Group(children),
	)
}