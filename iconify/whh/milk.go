package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.857 1024h-960q-13 0-22.5-9.5T.857 992V416q0-78 8-86l120-120V32q0-13 9.5-22.5t22.5-9.5h640q13 0 22.5 9.5t9.5 22.5v160l160 128q15 15 22.5 33.5t8.5 29.5t1 33v576q0 13-9.5 22.5t-22.5 9.5zm-928-64h512V448h-512v512zm0-576h512l128-128h-512zm704-320h-576v128h576V64zm32 192l-128 128h288zm160 192h-320v512h320V448z"/>`),
		g.Group(children),
	)
}