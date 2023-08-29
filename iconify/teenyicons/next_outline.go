package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m1.5 2.5l.29-.407A.5.5 0 0 0 1 2.5h.5Zm0 10H1a.5.5 0 0 0 .79.407L1.5 12.5Zm7-5l.29.407a.5.5 0 0 0 0-.814L8.5 7.5ZM1 2.5v10h1v-10H1Zm.79 10.407l7-5l-.58-.814l-7 5l.58.814Zm7-5.814l-7-5l-.58.814l7 5l.58-.814ZM13 2v11h1V2h-1Z"/>`),
		g.Group(children),
	)
}