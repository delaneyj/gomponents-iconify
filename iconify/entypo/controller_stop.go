package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 4.995v9.808c0 .661-.536 1.197-1.196 1.197H4.997A.997.997 0 0 1 4 15.003V5.196C4 4.536 4.536 4 5.196 4h9.808c.55 0 .996.446.996.995z"/>`),
		g.Group(children),
	)
}