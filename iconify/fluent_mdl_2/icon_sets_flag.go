package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IconSetsFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 0h128v2048H256V0zm1782 576L512 1148V4l1526 572zM640 964l1034-388L640 188v776z"/>`),
		g.Group(children),
	)
}