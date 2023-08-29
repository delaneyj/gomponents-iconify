package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M88.447 38.528H11.554a2.024 2.024 0 0 0-2.024 2.024v18.896c0 1.118.907 2.024 2.024 2.024h76.892a2.024 2.024 0 0 0 2.023-2.024V40.552a2.023 2.023 0 0 0-2.022-2.024z"/>`),
		g.Group(children),
	)
}