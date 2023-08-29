package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CashOutLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18.667 11C20.55 10.721 22 9.046 22 7.02C22 4.8 20.259 3 18.111 3H5.89C3.74 3 2 4.8 2 7.02C2 9.046 3.449 10.721 5.333 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 6v7m0 0l2-2.333M12 13l-2-2.333"/><path d="M5 10c0-1.886 0-2.828.586-3.414C6.172 6 7.114 6 9 6h6c1.886 0 2.828 0 3.414.586C19 7.172 19 8.114 19 10v6c0 1.886 0 2.828-.586 3.414C17.828 20 16.886 20 15 20H9c-1.886 0-2.828 0-3.414-.586C5 18.828 5 17.886 5 16v-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 17h14"/></g>`),
		g.Group(children),
	)
}