package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v1280H0V384h2048zM143 512l881 441l881-441H143zm1777 1024V648l-896 447l-896-447v888h1792z"/>`),
		g.Group(children),
	)
}