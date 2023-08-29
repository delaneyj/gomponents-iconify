package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 9.25c0-.69.56-1.25 1.25-1.25h25.5c.69 0 1.25.56 1.25 1.25V14a1.25 1.25 0 1 1-2.5 0v-3.5H25.25v27h3.5a1.25 1.25 0 1 1 0 2.5h-9.5a1.25 1.25 0 1 1 0-2.5h3.5v-27H12.5V14a1.25 1.25 0 1 1-2.5 0V9.25Z"/>`),
		g.Group(children),
	)
}