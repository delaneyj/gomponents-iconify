package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chevronup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m577 32l426 434q21 21 21 51t-21 51l-51 51q-21 21-51 21t-51-21L512 281L174 619q-21 21-51 21t-51-21l-51-51Q0 547 0 517t21-51L449 32q32-32 64-32t64 32z"/>`),
		g.Group(children),
	)
}