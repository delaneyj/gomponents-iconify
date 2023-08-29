package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 0 4.472 10H8v-1h5.197c.183-.316.338-.65.462-1H8V9h5.917c.055-.325.083-.66.083-1H8V7h5.917a5.952 5.952 0 0 0-.258-1H8V5h5.197a6.018 6.018 0 0 0-.725-1H8V3h3.318A5.97 5.97 0 0 0 8 2Z"/>`),
		g.Group(children),
	)
}