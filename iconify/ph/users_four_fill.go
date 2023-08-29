package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230.4 219.19A8 8 0 0 1 224 232H32a8 8 0 0 1-6.4-12.8A67.88 67.88 0 0 1 53 197.51a40 40 0 1 1 53.93 0a67.42 67.42 0 0 1 21 14.29a67.42 67.42 0 0 1 21-14.29a40 40 0 1 1 53.93 0a67.85 67.85 0 0 1 27.54 21.68ZM27.2 126.4a8 8 0 0 0 11.2-1.6a52 52 0 0 1 83.2 0a8 8 0 0 0 12.8 0a52 52 0 0 1 83.2 0a8 8 0 0 0 12.8-9.61A67.85 67.85 0 0 0 203 93.51a40 40 0 1 0-53.93 0a67.42 67.42 0 0 0-21 14.29a67.42 67.42 0 0 0-21-14.29a40 40 0 1 0-53.93 0A67.88 67.88 0 0 0 25.6 115.2a8 8 0 0 0 1.6 11.2Z"/>`),
		g.Group(children),
	)
}