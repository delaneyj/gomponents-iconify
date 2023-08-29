package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.332 2.626A1.25 1.25 0 0 1 4.415 2h3.17c.447 0 .86.239 1.083.626l1.585 2.75a1.25 1.25 0 0 1 0 1.248l-1.585 2.75A1.25 1.25 0 0 1 7.585 10h-3.17a1.25 1.25 0 0 1-1.083-.626l-1.586-2.75a1.25 1.25 0 0 1 0-1.248l1.586-2.75ZM4.415 3a.25.25 0 0 0-.217.125l-1.585 2.75a.25.25 0 0 0 0 .25l1.585 2.75A.25.25 0 0 0 4.415 9h3.17a.25.25 0 0 0 .216-.125l1.586-2.75a.25.25 0 0 0 0-.25l-1.586-2.75A.25.25 0 0 0 7.585 3h-3.17Z"/>`),
		g.Group(children),
	)
}