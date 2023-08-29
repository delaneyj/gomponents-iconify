package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClubsFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2a5 5 0 0 0-4.488 2.797l-.103.225a4.998 4.998 0 0 0-.334 2.837l.027.14a5 5 0 0 0-3.091 9.009l.198.14a4.998 4.998 0 0 0 4.42.58l.174-.066l-.773 3.095A1 1 0 0 0 9 22h6l.113-.006a1 1 0 0 0 .857-1.237l-.774-3.095l.174.065A5 5 0 1 0 16.897 8l.028-.14A4.997 4.997 0 0 0 12 2z"/></g>`),
		g.Group(children),
	)
}