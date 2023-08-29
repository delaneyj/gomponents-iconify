package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Command(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M9.455 14.91h5.454V9.454a5.455 5.455 0 1 0-5.454 5.454Zm0 18.18h5.454v5.455a5.455 5.455 0 1 1-5.454-5.454Z"/><path stroke-linecap="round" d="M14.909 14.909h18.182v18.182H14.909z"/><path d="M38.545 14.91h-5.454V9.454a5.455 5.455 0 1 1 5.454 5.454Zm0 18.18a5.455 5.455 0 1 1-5.454 5.455v-5.454h5.454Z"/></g>`),
		g.Group(children),
	)
}