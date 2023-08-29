package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosGridView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M192 192h128v128H192z" fill="currentColor"/><path d="M64 64v384h384V64H64zm352 128h-80v128h80v16h-80v80h-16v-80H192v80h-16v-80H96v-16h80V192H96v-16h80V96h16v80h128V96h16v80h80v16z" fill="currentColor"/>`),
		g.Group(children),
	)
}