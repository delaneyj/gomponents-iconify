package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CategoryNewEach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 10h-5v2h5v6h-7v2h3v2.142a4 4 0 1 0 2 0V20h2a2.003 2.003 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm-1 16a2 2 0 1 1-2-2a2.003 2.003 0 0 1 2 2zM19 6h-5v2h5v6h-7v2h3v6.142a4 4 0 1 0 2 0V16h2a2.002 2.002 0 0 0 2-2V8a2.002 2.002 0 0 0-2-2zm-1 20a2 2 0 1 1-2-2a2.003 2.003 0 0 1 2 2zM9 2H3a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h2v10.142a4 4 0 1 0 2 0V12h2a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2zM8 26a2 2 0 1 1-2-2a2.002 2.002 0 0 1 2 2zM3 10V4h6l.002 6z"/>`),
		g.Group(children),
	)
}