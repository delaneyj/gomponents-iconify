package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligionIslamReligionIslamMoonCrescentMuslimCultureStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.21 7a5.5 5.5 0 0 1 5.47-5.5a6.5 6.5 0 1 0 0 11A5.5 5.5 0 0 1 5.21 7Z"/><path d="m10.57 4.04l.91 1.81h1.81l-1.36 1.4l.43 2l-1.79-1l-1.71 1l.36-2l-1.36-1.4h1.81l.9-1.81z"/></g>`),
		g.Group(children),
	)
}