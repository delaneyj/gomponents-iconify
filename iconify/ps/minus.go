package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M363 213q21 0 21-21t-21-21H21q-21 0-21 21t21 21h342z"/>`),
		g.Group(children),
	)
}