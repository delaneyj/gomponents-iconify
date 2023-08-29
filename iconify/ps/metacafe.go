package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metacafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m421 462l-68-49l-68 49l26-79l-67-49h83l26-79l26 79h83l-67 49zm-129-73l-99-71h122l38-116q32 99 38 116h65L353 2l-83 254H2l269 196q4-12 8.5-25t8-24.5T292 389z"/>`),
		g.Group(children),
	)
}