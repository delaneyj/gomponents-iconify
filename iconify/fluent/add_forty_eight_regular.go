package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.75 5.5c.69 0 1.25.56 1.25 1.25V22.5h15.75a1.25 1.25 0 1 1 0 2.5H25v15.75a1.25 1.25 0 1 1-2.5 0V25H6.75a1.25 1.25 0 1 1 0-2.5H22.5V6.75c0-.69.56-1.25 1.25-1.25Z"/>`),
		g.Group(children),
	)
}