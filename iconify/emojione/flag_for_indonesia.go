package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForIndonesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M31.8 62c16.6 0 30-13.4 30-30h-60c0 16.6 13.4 30 30 30"/><path fill="#ed4c5c" d="M31.8 2c-16.6 0-30 13.4-30 30h60c0-16.6-13.4-30-30-30"/>`),
		g.Group(children),
	)
}