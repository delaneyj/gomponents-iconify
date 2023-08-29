package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassTiltedLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M5 13a8 8 0 1 0 16 0a8 8 0 0 0-16 0Zm12.348-4.268c.552.957.419 2.068-.299 2.482c-.717.414-1.747-.025-2.299-.982c-.552-.957-.418-2.068.299-2.482c.718-.414 1.747.025 2.3.982Z"/><path d="M2 13c0 6.075 4.925 11 11 11c2.295 0 4.426-.703 6.19-1.905a3.747 3.747 0 0 0 1.005 3.483l3.182 3.182a3.75 3.75 0 0 0 5.303-5.303l-3.182-3.182a3.747 3.747 0 0 0-3.454-1.012A10.95 10.95 0 0 0 24 13c0-6.075-4.925-11-11-11S2 6.925 2 13Zm20 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/></g>`),
		g.Group(children),
	)
}