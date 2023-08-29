package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plurk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M289 4H6v461h117V338h283V4H289zm0 218H123V120h166v102z"/>`),
		g.Group(children),
	)
}