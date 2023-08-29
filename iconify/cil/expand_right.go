package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M40 19h32v480H40zm352 480h128V19H392Zm32-448h64v416h-64ZM120 419h37.86l169.364-159.923L159.921 99H120Zm32-283.291l128.776 123.215L152 380.522Z"/>`),
		g.Group(children),
	)
}