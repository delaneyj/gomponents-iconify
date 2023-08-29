package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m259 733l90-90l701 701l-701 701l-90-90l611-611l-611-611zM1789 93l-611 611l611 611l-90 90l-701-701L1699 3l90 90z"/>`),
		g.Group(children),
	)
}