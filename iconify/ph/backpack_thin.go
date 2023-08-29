package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 44.17V32a20 20 0 0 0-20-20h-32a20 20 0 0 0-20 20v12.17A52.05 52.05 0 0 0 44 96v120a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V96a52.05 52.05 0 0 0-48-51.83ZM112 20h32a12 12 0 0 1 12 12v12h-56V32a12 12 0 0 1 12-12Zm60 144H84v-12a12 12 0 0 1 12-12h64a12 12 0 0 1 12 12Zm-88 8h56v12a4 4 0 0 0 8 0v-12h24v48H84Zm120 44a4 4 0 0 1-4 4h-20v-68a20 20 0 0 0-20-20H96a20 20 0 0 0-20 20v68H56a4 4 0 0 1-4-4V96a44.05 44.05 0 0 1 44-44h64a44.05 44.05 0 0 1 44 44ZM148 88a4 4 0 0 1-4 4h-32a4 4 0 0 1 0-8h32a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}