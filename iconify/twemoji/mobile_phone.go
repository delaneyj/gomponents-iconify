package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobilePhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#31373D" d="M11 36s-4 0-4-4V4s0-4 4-4h14s4 0 4 4v28s0 4-4 4H11z"/><path fill="#55ACEE" d="M9 5h18v26H9z"/>`),
		g.Group(children),
	)
}