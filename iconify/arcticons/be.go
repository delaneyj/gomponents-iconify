package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Be(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.954 33.135c-2.828-4.583-1.406-10.59 3.176-13.418c4.583-2.828 10.59-1.406 13.418 3.176c2.828 4.581 1.407 10.589-3.175 13.417c-4.582 2.828-10.589 1.407-13.418-3.174"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.625 36.457a9.75 9.75 0 1 1 4.875-8.443m-19.498 0H43.5m-38.997-.001V10.235"/>`),
		g.Group(children),
	)
}