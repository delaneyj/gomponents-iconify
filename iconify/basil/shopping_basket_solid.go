package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBasketSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.65 4.375a.75.75 0 0 0-1.3-.75L6.69 6.5H5.5a2 2 0 0 0-1.058 3.697a.362.362 0 0 0 .193.053h14.73a.362.362 0 0 0 .193-.053A2 2 0 0 0 18.5 6.5h-1.19l-1.66-2.875a.75.75 0 1 0-1.3.75L15.578 6.5H8.423L9.65 4.375Z"/><path fill="currentColor" fill-rule="evenodd" d="M19.335 11.985a.2.2 0 0 0-.197-.235H4.862a.2.2 0 0 0-.197.235l1 5.698a2.773 2.773 0 0 0 2.35 2.267c2.644.368 5.326.368 7.97 0a2.773 2.773 0 0 0 2.35-2.267l1-5.698ZM10.75 14a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0v-2Zm3.25-.75a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}