package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 10.286l6.857 6.857L12 24.001l-12-12l12-12l3 3l-1.714 1.714L12 3.429L3.429 12L12 20.571l3.446-3.446L10.285 12zM22.286 0l12 12l-12 12l-3-3L21 19.286l1.286 1.286l8.571-8.571l-8.571-8.571l-3.446 3.446l5.161 5.125l-1.714 1.714l-6.857-6.857z"/>`),
		g.Group(children),
	)
}