package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21v-2.125l5.3-5.3l2.125 2.125l-5.3 5.3H12Zm-9-5v-2h7v2H3Zm17.125-1L18 12.875l.725-.725q.275-.275.7-.275t.7.275l.725.725q.275.275.275.7t-.275.7l-.725.725ZM3 12v-2h11v2H3Zm0-4V6h11v2H3Z"/>`),
		g.Group(children),
	)
}