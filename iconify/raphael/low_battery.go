package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LowBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.27 13.5h-1.103v-1.416a1 1 0 0 0-1-1H5.25a1 1 0 0 0-1 1v7.832a1 1 0 0 0 1 1h19.917a1 1 0 0 0 1-1V18.5h1.104c.266 0 .48-.448.48-1v-3c0-.552-.214-1-.48-1zm-3.103 5.416H6.25v-5.832h17.917v5.832zm-15-4.832H7.25v3.832h1.917v-3.832z"/>`),
		g.Group(children),
	)
}