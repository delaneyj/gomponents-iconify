package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64v384h480V64Zm304 32v320H192V96ZM48 96h112v320H48Zm416 320H352V96h112Z"/>`),
		g.Group(children),
	)
}