package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.415 2c-.447 0-.86.239-1.083.626l-1.586 2.75a1.25 1.25 0 0 0 0 1.248l1.586 2.75A1.25 1.25 0 0 0 4.415 10h3.17c.447 0 .86-.239 1.083-.626l1.585-2.75a1.25 1.25 0 0 0 0-1.248l-1.585-2.75A1.25 1.25 0 0 0 7.585 2h-3.17Z"/>`),
		g.Group(children),
	)
}