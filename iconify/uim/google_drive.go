package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.334 14.887H22l-6.666-11.55H8.667l6.667 11.55z" opacity=".25"/><path fill="currentColor" d="M8.667 3.338L2 14.888l3.334 5.774L12 9.112z"/><path fill="currentColor" d="m8.667 14.887l-3.333 5.775h13.333L22 14.887z" opacity=".5"/>`),
		g.Group(children),
	)
}