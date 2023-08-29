package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationEffect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 40.833a11.955 11.955 0 0 0 8 3.056c6.627 0 12-5.373 12-12c0-5.301-3.437-9.8-8.204-11.387"/><path d="M27.171 27.5c.535 1.359.829 2.84.829 4.39c0 6.627-5.373 12-12 12c-6.628 0-12-5.373-12-12c0-5.316 3.455-9.824 8.242-11.4"/><path d="M24 27.89c6.627 0 12-5.373 12-12c0-6.628-5.373-12-12-12s-12 5.372-12 12c0 6.627 5.373 12 12 12Z"/></g>`),
		g.Group(children),
	)
}