package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.71 21.29l-1-1l-6-6l-6-6l-6-6l-1-1a1 1 0 0 0-1.42 1.42l.71.7V21a1 1 0 0 0 1 1h17.59l.7.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM4 5.41L6.59 8H4ZM8 20H4v-4h4Zm0-6H4v-4h4Zm2-2.59L12.59 14H10ZM14 20h-4v-4h4Zm2 0v-2.59L18.59 20ZM8.67 4H14v5a1 1 0 0 0 1 1h5v5.33a1 1 0 1 0 2 0V3a1 1 0 0 0-1-1H8.67a1 1 0 0 0 0 2ZM16 4h4v4h-4Z"/>`),
		g.Group(children),
	)
}