package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chevrondown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m577 608l426-434q21-21 21-51t-21-51l-51-51Q931 0 901 0t-51 21L512 359L174 21Q153 0 123 0T72 21L21 72Q0 93 0 123t21 51l428 434q32 32 64 32t64-32z"/>`),
		g.Group(children),
	)
}