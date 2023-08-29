package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="none" d="M128 128H0V0h128v128z"/><circle cx="64" cy="64" r="32" fill="#40c0e7"/>`),
		g.Group(children),
	)
}