package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePptTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.859 2.878l12.57-1.796a.5.5 0 0 1 .571.495v20.847a.5.5 0 0 1-.57.495L2.858 21.123a1 1 0 0 1-.859-.99V3.868a1 1 0 0 1 .859-.99ZM4 4.735v14.53l10 1.43V3.305L4 4.735ZM17 19h3V5h-3V3h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-4v-2ZM5 8h8v6H7v2H5V8Zm2 2v2h4v-2H7Z"/>`),
		g.Group(children),
	)
}