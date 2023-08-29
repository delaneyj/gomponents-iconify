package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnticlockwiseTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.414 6l1.829 1.828l-1.415 1.415L9.586 5L13.828.757l1.415 1.415L13.414 4H16a5 5 0 0 1 5 5v4h-2V9a3 3 0 0 0-3-3h-2.586ZM15 11v10a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V11a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm-2 1H5v8h8v-8Z"/>`),
		g.Group(children),
	)
}