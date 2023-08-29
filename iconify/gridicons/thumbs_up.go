package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.7 22H2v-9h2l2.7 9zM19.999 9H14V5a3 3 0 0 0-3-3h-1v4L7.1 9.625A5.02 5.02 0 0 0 6 12.76V14l2.1 7h8.337a4 4 0 0 0 3.881-3.03l1.621-6.485A2 2 0 0 0 19.999 9z"/>`),
		g.Group(children),
	)
}