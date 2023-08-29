package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastRewindOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.5 18l-9-6l9-6v12Zm-10 0l-9-6l9-6v12Zm-2-6Zm10 0Zm-10 2.25v-4.5L6.1 12l3.4 2.25Zm10 0v-4.5L16.1 12l3.4 2.25Z"/>`),
		g.Group(children),
	)
}