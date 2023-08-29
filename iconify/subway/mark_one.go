package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M213.3 0h-128C38.2 0 0 38.2 0 85.3v128L298.7 512L512 298.7L213.3 0zm-128 128c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7S128 61.8 128 85.3S108.9 128 85.3 128zm85.4 192L320 170.7l42.7 42.7l-149.4 149.3l-42.6-42.7zm85.3 85.3L405.3 256l42.7 42.7L298.7 448L256 405.3z"/>`),
		g.Group(children),
	)
}