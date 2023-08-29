package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableTennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 14c1.4 0 2.5 1.1 2.5 2.5S19.9 19 18.5 19S16 17.9 16 16.5s1.1-2.5 2.5-2.5M7 15s1 1 1 2v3.5c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5V17c0-1 1-2 1-2H7m1-1h3s5 0 5-5s-4-7-6.5-7S3 4 3 9s5 5 5 5Z"/>`),
		g.Group(children),
	)
}