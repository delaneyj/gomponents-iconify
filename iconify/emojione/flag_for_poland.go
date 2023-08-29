package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPoland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M32 2c16.6 0 30 13.4 30 30H2C2 15.4 15.4 2 32 2z"/><path fill="#ed4c5c" d="M32 62C15.4 62 2 48.6 2 32h60c0 16.6-13.4 30-30 30"/>`),
		g.Group(children),
	)
}