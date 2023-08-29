package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M338 0H103.3v426.7h384V149.3H338V0zm21.3 0v128h128L359.3 0zM60.7 85.3H18V512h384v-42.7H60.7v-384z"/>`),
		g.Group(children),
	)
}