package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 16c1.403-.234 3.637-.293 5.945.243M16 21c-1.704-2.768-4.427-4.148-7.055-4.757m0 0C10.895 13.985 14.558 12 21 12h1M8.5 7C8 7 7 7.3 7 8.5S8 10 8.5 10S10 9.7 10 8.5S9 7 8.5 7z"/><circle cx="12" cy="12" r="10"/></g>`),
		g.Group(children),
	)
}