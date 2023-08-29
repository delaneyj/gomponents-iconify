package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivisionCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopDivisionCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" transform="translate(3 3)"><path d="M5 11a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2H5Z"/><circle cx="10" cy="5.5" r="1.5"/><circle cx="10" cy="14.5" r="1.5"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopDivisionCircleFilled0)"/></g>`),
		g.Group(children),
	)
}