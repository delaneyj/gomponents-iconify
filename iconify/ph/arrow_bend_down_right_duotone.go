package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDownRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 152l-48 48v-96Z" opacity=".2"/><path d="m229.66 146.34l-48-48A8 8 0 0 0 168 104v40h-40a88.1 88.1 0 0 1-88-88a8 8 0 0 0-16 0a104.11 104.11 0 0 0 104 104h40v40a8 8 0 0 0 13.66 5.66l48-48a8 8 0 0 0 0-11.32ZM184 180.69v-57.38L212.69 152Z"/></g>`),
		g.Group(children),
	)
}