package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronSmallLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.141 13.418a.695.695 0 0 1 0 .978a.68.68 0 0 1-.969 0l-3.83-3.908a.697.697 0 0 1 0-.979l3.83-3.908a.68.68 0 0 1 .969 0a.695.695 0 0 1 0 .978L9 10l3.141 3.418z"/>`),
		g.Group(children),
	)
}