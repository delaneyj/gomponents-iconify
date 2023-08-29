package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M41.355 34.153L29.522 8.776a6 6 0 0 0-10.876 5.072L30.48 39.224a6 6 0 1 0 10.876-5.071Z"/><path stroke-linecap="round" d="M23.438 26.536L17.52 39.224a6 6 0 0 1-7.974 2.902v0a6 6 0 0 1-2.902-7.973L18.374 9"/><circle cx="12.083" cy="36.688" r="6" transform="rotate(25 12.083 36.688)"/></g>`),
		g.Group(children),
	)
}