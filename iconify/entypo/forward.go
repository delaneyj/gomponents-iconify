package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 11.874v4.357l7-6.69l-7-6.572v3.983c-8.775 0-11 9.732-11 9.732c2.484-4.388 6.237-4.81 11-4.81z"/>`),
		g.Group(children),
	)
}