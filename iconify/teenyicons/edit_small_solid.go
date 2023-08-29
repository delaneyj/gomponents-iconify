package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8.854 4.146a.5.5 0 0 0-.708 0L4 8.293V10.5a.5.5 0 0 0 .5.5h2.207l4.147-4.146a.5.5 0 0 0 0-.708l-2-2Z"/>`),
		g.Group(children),
	)
}