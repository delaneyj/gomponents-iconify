package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64v384h480V64Zm448 176H352V96h112Zm-272 0V96h128v144Zm128 32v144H192V272ZM160 96v144H48V96ZM48 272h112v144H48Zm304 144V272h112v144Z"/>`),
		g.Group(children),
	)
}