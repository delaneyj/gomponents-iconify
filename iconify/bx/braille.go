package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braille(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="4" cy="7" r="2" fill="currentColor"/><circle cx="9" cy="12" r="2" fill="currentColor"/><circle cx="15" cy="7" r="2" fill="currentColor"/><circle cx="15" cy="12" r="2" fill="currentColor"/><circle cx="15" cy="17" r="2" fill="currentColor"/><circle cx="20" cy="7" r="2" fill="currentColor"/><circle cx="4" cy="17" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}