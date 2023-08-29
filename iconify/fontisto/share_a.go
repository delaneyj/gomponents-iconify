package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.408 22.528C2.281 15.77 6.08 10.995 11.624 8.434c1.678-.752 3.633-1.353 5.673-1.709l.151-.022c.462 0 .464-.014.464-3.352V0l11.446 11.446l-11.446 11.446V16.19H16.52a24.855 24.855 0 0 0-10.51 2.58l.145-.065a20.31 20.31 0 0 0-4.767 3.825l-.013.015l-1.374 1.454l.408-1.472z"/>`),
		g.Group(children),
	)
}