package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinybit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 44.5l17.799-10.276V13.8L24 24.076V44.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24.076L6.201 13.8L23.977 3.5l17.822 10.3L24 24.076z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 44.5L6.201 34.224V13.8L24 24.076V44.5zm0-20.424l-4.603-2.658l4.597-2.663l4.609 2.663L24 24.076z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 29.358L19.397 26.7v-5.282L24 24.076v5.282zm0 0l4.603-2.658v-5.282L24 24.076v5.282z"/>`),
		g.Group(children),
	)
}