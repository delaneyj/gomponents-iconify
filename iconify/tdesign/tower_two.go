package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .586l2.973 2.973L15.936 18H17v5H7v-5h1.064l.963-14.44L12 .586ZM10.069 18h3.862l-.904-13.558L12 3.414l-1.027 1.028L10.069 18ZM9 20v1h6v-1H9Z"/>`),
		g.Group(children),
	)
}