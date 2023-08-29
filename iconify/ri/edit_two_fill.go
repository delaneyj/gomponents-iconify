package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.243 18.996H21v2H3v-4.242l9.9-9.9l4.242 4.243l-7.9 7.9Zm5.07-13.556l2.122-2.121a1 1 0 0 1 1.414 0l2.829 2.828a1 1 0 0 1 0 1.414l-2.122 2.122l-4.242-4.243Z"/>`),
		g.Group(children),
	)
}