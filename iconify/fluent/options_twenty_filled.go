package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionsTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.25 5h-2.364a2.501 2.501 0 0 0-4.771 0H2.75a.75.75 0 0 0 0 1.5h7.364a2.501 2.501 0 0 0 4.771 0h2.365a.75.75 0 0 0 0-1.5Zm-14.5 8.5a.75.75 0 0 0 0 1.5h2.364a2.501 2.501 0 0 0 4.772 0h7.364a.75.75 0 0 0 0-1.5H9.886a2.501 2.501 0 0 0-4.772 0H2.75Z"/>`),
		g.Group(children),
	)
}