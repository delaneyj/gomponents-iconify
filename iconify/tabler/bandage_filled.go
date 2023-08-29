package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandageFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M20.207 3.793a5.95 5.95 0 0 1 .179 8.228l-.179.186l-8 8a5.95 5.95 0 0 1-8.593-8.228l.179-.186l8-8a5.95 5.95 0 0 1 8.414 0zM12 13a1 1 0 0 0-1 1l.007.127A1 1 0 0 0 13 14.01l-.007-.127A1 1 0 0 0 12 13zm2-2a1 1 0 0 0-1 1l.007.127A1 1 0 0 0 15 12.01l-.007-.127A1 1 0 0 0 14 11zm-4 0a1 1 0 0 0-1 1l.007.127A1 1 0 0 0 11 12.01l-.007-.127A1 1 0 0 0 10 11zm2-2a1 1 0 0 0-1 1l.007.127A1 1 0 0 0 13 10.01l-.007-.127A1 1 0 0 0 12 9z"/></g>`),
		g.Group(children),
	)
}