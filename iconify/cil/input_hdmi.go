package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputHdmi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 16v104H80v131.492l72 157.091V496h208v-87.417l72-157.091V120h-32V16Zm32 32h224v72h-32V80h-32v40h-32V80h-32v40h-32V80h-32v40h-32Zm256 196.508L328 401.6V464H184v-62.4l-72-157.092V152h288Z"/>`),
		g.Group(children),
	)
}