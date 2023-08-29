package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCodesandbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 7.5v9l-4 2.25L12 21l-4-2.25l-4-2.25v-9l4-2.25L12 3l4 2.25zM12 12l4-2.25l4-2.25M12 12v9m0-9L8 9.75L4 7.5"/><path d="m20 12l-4 2v4.75M4 12l4 2v4.75m0-13.5l4 2.25l4-2.25"/></g>`),
		g.Group(children),
	)
}