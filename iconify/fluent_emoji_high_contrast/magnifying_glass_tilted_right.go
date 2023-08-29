package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassTiltedRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.779 21a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm4.27-9.786c-.717.414-1.747-.025-2.299-.982c-.552-.957-.418-2.068.299-2.482c.718-.414 1.747.025 2.3.982c.551.957.418 2.068-.3 2.482Z"/><path d="M18.779 24c6.075 0 11-4.925 11-11s-4.925-11-11-11s-11 4.925-11 11c0 2.327.722 4.485 1.955 6.263a3.747 3.747 0 0 0-3.454 1.012l-3.182 3.182a3.75 3.75 0 1 0 5.304 5.303l3.182-3.182a3.747 3.747 0 0 0 1.005-3.483A10.949 10.949 0 0 0 18.78 24Zm0-2a9 9 0 1 1 0-18a9 9 0 0 1 0 18Z"/></g>`),
		g.Group(children),
	)
}