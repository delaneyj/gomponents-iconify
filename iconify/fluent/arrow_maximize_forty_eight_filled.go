package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M40.5 24a1.5 1.5 0 0 1-1.5-1.5V11.121L11.121 39H22.5a1.5 1.5 0 0 1 0 3h-15A1.5 1.5 0 0 1 6 40.5v-15a1.5 1.5 0 0 1 3 0v11.379L36.879 9H25.5a1.5 1.5 0 0 1 0-3h15A1.5 1.5 0 0 1 42 7.5v15a1.5 1.5 0 0 1-1.5 1.5Z"/>`),
		g.Group(children),
	)
}