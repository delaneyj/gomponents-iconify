package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipRightDesignFlipReflectRightHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m0 3v3m-2 5h1a1 1 0 0 0 1-1v-1m-9 2h-3a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h3m4 13H7V.5h1.5"/>`),
		g.Group(children),
	)
}