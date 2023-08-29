package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M37.826 4A19.629 19.629 0 0 1 44 18.316C44 29.187 35.187 38 24.316 38A19.629 19.629 0 0 1 10 31.826"/><path d="M24 32c7.732 0 14-6.268 14-14S31.732 4 24 4s-14 6.268-14 14s6.268 14 14 14Z" clip-rule="evenodd"/><path d="M24 38v6m-6 0h12"/></g>`),
		g.Group(children),
	)
}