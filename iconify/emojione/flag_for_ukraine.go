package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForUkraine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffe62e" d="M32 62c16.6 0 30-13.4 30-30H2c0 16.6 13.4 30 30 30z"/><path fill="#428bc1" d="M32 2C15.4 2 2 15.4 2 32h60C62 15.4 48.6 2 32 2z"/>`),
		g.Group(children),
	)
}