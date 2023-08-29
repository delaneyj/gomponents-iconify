package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="13" cy="27" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 21c0-5.5 2-16 11-16c4.5 0 12.72 2.332 9.536 15.561c-.28 1.166-1.382 3.414-1.382 3.414l-1.604 3.602c-1.659 3.341-3.802 11.877 2.986 15.034c2.058.957 6.481.339 7.464-4.01"/></g>`),
		g.Group(children),
	)
}