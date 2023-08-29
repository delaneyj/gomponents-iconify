package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64v384h480V64Zm448 176H192V96h272Zm-272 32h120v144H192ZM48 96h112v320H48Zm296 320V272h120v144Z"/>`),
		g.Group(children),
	)
}