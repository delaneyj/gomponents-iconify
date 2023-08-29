package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelmetOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 11c-9.389 0-17 7.815-17 17.454V35h34v-6.546C41 18.814 33.389 11 24 11ZM4 35h40v6H4z"/><path d="M21 6h6v18h-6z"/></g>`),
		g.Group(children),
	)
}