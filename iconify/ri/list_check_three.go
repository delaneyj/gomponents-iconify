package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 6v3H5V6h3ZM3 4v7h7V4H3Zm10 0h8v2h-8V4Zm0 7h8v2h-8v-2Zm0 7h8v2h-8v-2Zm-2.293-1.793l-1.414-1.414L6 18.086l-1.793-1.793l-1.414 1.414L6 20.914l4.707-4.707Z"/>`),
		g.Group(children),
	)
}