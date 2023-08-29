package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotSingle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.8 10a2.2 2.2 0 0 0 4.4 0a2.2 2.2 0 0 0-4.4 0z"/>`),
		g.Group(children),
	)
}