package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M480 256l-96-96v64h-96v-96h64l-96-96-96 96h64v96h-96v-64l-96 96 96 96v-64h96v96h-64l96 96 96-96h-64v-96h96v64z" fill="currentColor"/>`),
		g.Group(children),
	)
}