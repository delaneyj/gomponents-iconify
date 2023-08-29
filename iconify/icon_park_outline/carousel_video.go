package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M11 7h26v34H11zm-7 4h7v26H4zm33 0h7v26h-7z"/><path d="m22 20l6 4l-6 4v-8Z"/></g>`),
		g.Group(children),
	)
}