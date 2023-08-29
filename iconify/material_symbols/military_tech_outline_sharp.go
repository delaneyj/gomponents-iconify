package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilitaryTechOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.9 22l1.2-3.8L7 16h3.8l.7-2.3L7 11.05V2h10v9.05l-4.5 2.65l.7 2.3H17l-3.1 2.2l1.2 3.8l-3.1-2.35L8.9 22ZM9 4v5.85l2 1.2V4H9Zm6 0h-2v7.05l2-1.2V4Zm-3 3.825Zm-1-.3Zm2 0Z"/>`),
		g.Group(children),
	)
}