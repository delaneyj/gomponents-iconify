package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.5 5A1.5 1.5 0 0 1 25 6.5V22h15.5a1.5 1.5 0 0 1 0 3H25v15.5a1.5 1.5 0 0 1-3 0V25H6.5a1.5 1.5 0 0 1 0-3H22V6.5A1.5 1.5 0 0 1 23.5 5Z"/>`),
		g.Group(children),
	)
}