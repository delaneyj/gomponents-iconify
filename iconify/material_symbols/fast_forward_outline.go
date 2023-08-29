package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 18V6l9 6l-9 6Zm10 0V6l9 6l-9 6Zm-8-6Zm10 0Zm-10 2.25L7.9 12L4.5 9.75v4.5Zm10 0L17.9 12l-3.4-2.25v4.5Z"/>`),
		g.Group(children),
	)
}