package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewStream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 64v384h480V64Zm448 32v144H48V96ZM48 416V272h416v144Z"/>`),
		g.Group(children),
	)
}