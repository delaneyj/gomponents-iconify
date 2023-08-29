package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.604 3.332C12.99 4 12.075 2.833 10 1C7.925 2.833 7.01 4 2.396 3.332C-.063 15.58 10 19 10 19s10.063-3.42 7.604-15.668zm-5.131 9.977L10 12.009l-2.472 1.3L8 10.556l-2-1.95l2.764-.401L10 5.7l1.236 2.505L14 8.606l-2 1.949l.473 2.754z"/>`),
		g.Group(children),
	)
}