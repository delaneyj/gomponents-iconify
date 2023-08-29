package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M11 9.472A4.49 4.49 0 0 0 9.682 6.29A4.481 4.481 0 0 0 7 5m4 4.472a4.5 4.5 0 0 1-9 0m9 0H2m0 0A4.49 4.49 0 0 1 3.318 6.29m11.439.831A3 3 0 1 0 19 2.88a3 3 0 0 0-4.243 4.242Zm2.624 13.784l3.524-3.524a3.738 3.738 0 0 0-5.286-5.286l-3.524 3.524A3.738 3.738 0 0 0 14 21.927"/><path d="M14 14s.281 1.46 1.911 3.09C17.541 18.718 19 19 19 19"/></g>`),
		g.Group(children),
	)
}