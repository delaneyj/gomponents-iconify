package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M322 0v128h128L322 0zm-21.3 0H66v288l64-64l85.3 85.3l85.3-85.3l85.3 85.3l64-64v-96H300.7V0zm0 309.3l-85.3 85.3l-85.4-85.3l-64 64V512h384V330.7l-64 64l-85.3-85.4z"/>`),
		g.Group(children),
	)
}