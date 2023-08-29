package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CashOutLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18.667 12C20.55 11.721 22 10.046 22 8.02C22 5.8 20.259 4 18.111 4H5.89C3.74 4 2 5.8 2 8.02c0 2.026 1.449 3.701 3.333 3.98" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v7m0 0l2-2.333M12 14l-2-2.333"/><path d="M5 11c0-1.886 0-2.828.586-3.414C6.172 7 7.114 7 9 7h6c1.886 0 2.828 0 3.414.586C19 8.172 19 9.114 19 11v6c0 1.886 0 2.828-.586 3.414C17.828 21 16.886 21 15 21H9c-1.886 0-2.828 0-3.414-.586C5 19.828 5 18.886 5 17v-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 18h14" opacity=".5"/></g>`),
		g.Group(children),
	)
}