package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M213.3 0h-128C38.2 0 0 38.2 0 85.3v128L298.7 512L512 298.7L213.3 0zm-128 128c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7S128 61.8 128 85.3S108.9 128 85.3 128z"/>`),
		g.Group(children),
	)
}