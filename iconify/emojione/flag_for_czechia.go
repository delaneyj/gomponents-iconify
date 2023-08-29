package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCzechia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M10.8 10.8V32H62C62 15.4 48.6 2 32 2c-8.3 0-15.8 3.4-21.2 8.8"/><path fill="#ed4c5c" d="M10.8 32v21.2C16.2 58.6 23.7 62 32 62c16.6 0 30-13.4 30-30H10.8"/><path fill="#428bc1" d="M10.8 10.8C5.4 16.2 2 23.7 2 32s3.4 15.8 8.8 21.2L32 32L10.8 10.8z"/>`),
		g.Group(children),
	)
}