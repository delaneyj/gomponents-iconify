package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m10 2.5l2.446 5.335h5.47l-4.582 4.24l1.559 5.425L10 14.687L5.108 17.5l1.559-5.426l-4.584-4.239h5.471z"/>`),
		g.Group(children),
	)
}