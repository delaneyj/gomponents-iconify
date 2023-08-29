package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedRectangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 6H8c-2.8 0-5 2.2-5 5v10c0 2.8 2.2 5 5 5h16c2.8 0 5-2.2 5-5V11c0-2.8-2.2-5-5-5z"/>`),
		g.Group(children),
	)
}