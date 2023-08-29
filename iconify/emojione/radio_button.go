package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="31" r="30" fill="#94989b"/><circle cx="32" cy="32" r="15" fill="#626a6d"/><circle cx="32" cy="30.4" r="15" fill="#d9e3e8"/>`),
		g.Group(children),
	)
}