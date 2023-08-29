package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 14a4 4 0 1 1 4-4a4.012 4.012 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.006 2.006 0 0 0-2-2Z"/><path fill="currentColor" d="M24 2H8a2.002 2.002 0 0 0-2 2v24a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2ZM12 24v-6a.945.945 0 0 1 1-1h6a.945.945 0 0 1 1 1v6h-2v4h-4v-4Zm8 4v-2a2.006 2.006 0 0 0 2-2v-6a2.946 2.946 0 0 0-3-3h-6a2.946 2.946 0 0 0-3 3v6a2.006 2.006 0 0 0 2 2v2H8V4h16v24Z"/>`),
		g.Group(children),
	)
}