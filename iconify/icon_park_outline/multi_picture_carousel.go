package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiPictureCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="30" x="4" y="6" rx="2"/><path d="M20 42h8m6 0h2M4 42h2m36 0h2m-32 0h2"/></g>`),
		g.Group(children),
	)
}