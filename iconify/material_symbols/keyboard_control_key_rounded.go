package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardControlKeyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.825L7.1 12.7q-.275.275-.687.288T5.7 12.7q-.275-.275-.275-.7t.275-.7l5.6-5.6q.3-.3.7-.3t.7.3l5.6 5.6q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L12 7.825Z"/>`),
		g.Group(children),
	)
}