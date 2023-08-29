package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 44h-64a43.86 43.86 0 0 0-32 13.85A43.86 43.86 0 0 0 96 44H32a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h64a20 20 0 0 1 20 20a12 12 0 0 0 24 0a20 20 0 0 1 20-20h64a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20ZM96 188H36V68h60a20 20 0 0 1 20 20v104.81A43.79 43.79 0 0 0 96 188Zm124 0h-60a43.71 43.71 0 0 0-20 4.83V88a20 20 0 0 1 20-20h60Z"/>`),
		g.Group(children),
	)
}