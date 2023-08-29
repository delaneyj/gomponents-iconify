package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1027 128l555 1664h-135l-170-512H643l-170 512H338L893 128h134zm207 1024L960 330l-274 822h548zm814-768h-640l320-320l320 320z"/>`),
		g.Group(children),
	)
}