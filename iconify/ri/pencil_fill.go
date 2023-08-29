package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.9 6.854l4.242 4.243l-9.9 9.9H3v-4.243l9.9-9.9Zm1.414-1.414l2.121-2.121a1 1 0 0 1 1.414 0l2.829 2.828a1 1 0 0 1 0 1.414l-2.122 2.122l-4.242-4.243Z"/>`),
		g.Group(children),
	)
}