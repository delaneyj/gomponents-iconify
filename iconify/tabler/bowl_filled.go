package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M20 7H4a2 2 0 0 0-2 2v.5l.007.18c.134 1.806 2.169 5.275 3.928 6.771l.065.053V17a2 2 0 0 0 2 2h8l.15-.005A2 2 0 0 0 18 17v-.504l.017-.013C19.753 14.989 22 11.194 22 9.5V9a2 2 0 0 0-2-2z"/></g>`),
		g.Group(children),
	)
}