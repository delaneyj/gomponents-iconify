package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.678 14.262h6.089V24H.678v-9.738zm8.215 9.717h6.089V7.11H8.893v16.869zM17.234 0v24h6.089V0h-6.089z"/>`),
		g.Group(children),
	)
}