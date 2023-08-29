package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBangladesh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#75a843"/><circle cx="32" cy="32" r="30" fill="#699635"/><circle cx="26" cy="32" r="14.1" fill="#ed4c5c"/>`),
		g.Group(children),
	)
}