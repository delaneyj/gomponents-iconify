package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForJapan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#f5f5f5"/><circle cx="32" cy="32" r="12" fill="#ed4c5c"/>`),
		g.Group(children),
	)
}