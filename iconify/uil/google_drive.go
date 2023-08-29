package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.423 13.888l-6.09-10.55H8.668l6.09 10.55ZM8.09 4.338L2 14.888l3.334 5.774l6.089-10.55Zm1.733 10.55L6.49 20.662h12.178L22 14.887Z"/>`),
		g.Group(children),
	)
}