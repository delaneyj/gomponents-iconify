package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontColorA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1582 1664h-135l-170-512H643l-170 512H338L893 0h134l555 1664zm-348-640L960 203l-274 821h548z"/>`),
		g.Group(children),
	)
}