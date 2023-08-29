package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieEditRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3l2 4h3L7 3h2l2 4h3l-2-4h2l2 4h3l-2-4h3q.825 0 1.413.588T22 5v4H4v8h8v2H4Zm15.15-6.3l2.15 2.1l-4.875 4.9q-.15.15-.338.225T15.7 20h-1.2q-.2 0-.35-.15T14 19.5v-1.2q0-.2.088-.4t.212-.325l4.85-4.875Zm2.875 1.4L19.9 11.975l.7-.7q.3-.3.725-.3t.7.3l.7.725q.275.3.275.713t-.275.687l-.7.7Z"/>`),
		g.Group(children),
	)
}