package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3l2 4h3L7 3h2l2 4h3l-2-4h2l2 4h3l-2-4h3q.825 0 1.413.588T22 5v4H4v8h8v2H4Zm10 1v-2.125l5.15-5.175l2.15 2.1l-5.175 5.2H14Zm8.025-5.9L19.9 11.975l.7-.7q.3-.3.725-.3t.7.3l.7.725q.275.3.275.713t-.275.687l-.7.7Z"/>`),
		g.Group(children),
	)
}