package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Charging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.27 13.5h-1.103v-1.416a1 1 0 0 0-1-1H5.25a1 1 0 0 0-1 1v7.832a1 1 0 0 0 1 1h19.917a1 1 0 0 0 1-1V18.5h1.104c.266 0 .48-.448.48-1v-3c0-.552-.214-1-.48-1zm-8.745 3.342l-6.733 3.366l3.366-3.366l-1.683-1.684l6.733-3.366l-3.366 3.366l1.683 1.684z"/>`),
		g.Group(children),
	)
}