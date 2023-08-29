package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kahoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.557 18.87l2.747-17.513L16.174 0zM.696 2.348v19.078l4.035.14l-.035-6.679l2.487-2.4l2.626 9.078h3.565L10.087 9.722l4.957-5.444l-3.496-1.339L4.73 9.443V1.322zm18.295 17.86l-.99 2.331L20.12 24l2.088-1.235l-.887-2.556Z"/>`),
		g.Group(children),
	)
}