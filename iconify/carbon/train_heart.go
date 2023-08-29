package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.303 2a2.661 2.661 0 0 0-1.908.806l-.393.405l-.397-.405a2.661 2.661 0 0 0-3.816 0a2.8 2.8 0 0 0 0 3.896L25.002 11l4.209-4.298a2.8 2.8 0 0 0 0-3.896A2.661 2.661 0 0 0 27.303 2zM22 13v3H6v-6h9V8H6.184A2.995 2.995 0 0 1 9 6h6V4H9a5.006 5.006 0 0 0-5 5v12a4.99 4.99 0 0 0 3.582 4.77L5.769 30h2.176l1.714-4h8.682l1.714 4h2.176l-1.813-4.23A4.99 4.99 0 0 0 24 21v-8zm0 7h-3v2h2.816A2.995 2.995 0 0 1 19 24H9a2.995 2.995 0 0 1-2.816-2H9v-2H6v-2h16z"/>`),
		g.Group(children),
	)
}