package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M71.49 60.55a12 12 0 0 0-23 0l-36 120A12 12 0 0 0 24 196h72a12 12 0 0 0 11.49-15.45ZM40.13 172L60 105.76L79.87 172ZM212 74a54 54 0 1 0-54 54a54.06 54.06 0 0 0 54-54Zm-84 0a30 30 0 1 1 30 30a30 30 0 0 1-30-30Zm96 70h-88a12 12 0 0 0-12 12v52a12 12 0 0 0 12 12h88a12 12 0 0 0 12-12v-52a12 12 0 0 0-12-12Zm-12 52h-64v-28h64Z"/>`),
		g.Group(children),
	)
}