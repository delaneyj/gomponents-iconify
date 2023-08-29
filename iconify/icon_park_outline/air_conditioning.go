package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirConditioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="20" x="4" y="8" rx="2"/><path d="M12 20h24v8H12zm20-6h4M24 34v6m-8-4v2m16-2v2"/></g>`),
		g.Group(children),
	)
}