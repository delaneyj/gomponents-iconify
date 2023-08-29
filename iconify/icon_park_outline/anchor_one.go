package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M40 35c0-9.205-7.163-25-16-25S8 25.795 8 35"/><path d="M4 35h8v8H4zM4 6h8v8H4zm32 29h8v8h-8zm0-29h8v8h-8z"/><path stroke-linecap="round" d="M12 10h24"/></g>`),
		g.Group(children),
	)
}