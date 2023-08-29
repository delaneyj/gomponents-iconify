package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 2h-5A1.5 1.5 0 0 0 2 3.5v.514A2.25 2.25 0 0 1 4.486 6h3.028A2.25 2.25 0 0 1 10 4.014V3.5A1.5 1.5 0 0 0 8.5 2Zm1.25 3A1.25 1.25 0 0 1 11 6.25V8.5a1.5 1.5 0 0 1-1 1.415v.585a.5.5 0 0 1-1 0V10H3v.5a.5.5 0 0 1-1 0v-.585A1.5 1.5 0 0 1 1 8.5V6.25a1.25 1.25 0 0 1 2.5 0c0 .414.336.75.75.75h3.5a.75.75 0 0 0 .75-.75C8.5 5.56 9.06 5 9.75 5Z"/>`),
		g.Group(children),
	)
}