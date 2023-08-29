package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartDiamondOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16.175L13.175 12L9 7.825L4.825 12L9 16.175ZM8.3 18.3l-5.6-5.6q-.3-.3-.3-.7t.3-.7l5.6-5.6q.3-.3.7-.3t.7.3L15 11h6q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-6l-5.3 5.3q-.3.3-.7.3t-.7-.3ZM9 12Z"/>`),
		g.Group(children),
	)
}