package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="m24.5 30.5l7.96 7.371L40.5 30.5h-5V9.549c0-2.5-.561-3.049-3-3.049h-18l6.641 6H29.5v18h-5zm-8-18L8.52 5.16L.5 12.5h5v21.049c0 2.5.62 2.951 3 2.951h18.32l-6.32-6h-9v-18h5z"/>`),
		g.Group(children),
	)
}