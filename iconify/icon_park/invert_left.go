package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 15C10 22.2989 14.103 28.5832 20 31.4081C22.1347 32.4307 24.5046 33 27 33C36.3888 33 44 24.9411 44 15"/><path d="M18 20L10 15L4 23"/></g>`),
		g.Group(children),
	)
}