package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForColombia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M62 32H2c0 5.5 1.5 10.6 4 15h52c2.6-4.4 4-9.5 4-15"/><path fill="#ffe62e" d="M32 2C15.5 2 2 15.4 2 32h60C62 15.4 48.6 2 32 2z"/><path fill="#ed4c5c" d="M32 62c11.1 0 20.8-6 26-15H6c5.3 9 14.9 15 26 15"/>`),
		g.Group(children),
	)
}