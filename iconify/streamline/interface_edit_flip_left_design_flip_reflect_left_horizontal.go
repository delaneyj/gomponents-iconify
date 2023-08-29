package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipLeftDesignFlipReflectLeftHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h-1a1 1 0 0 0-1 1v1m0 3v3m2 5h-1a1 1 0 0 1-1-1v-1m9 2h3a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1h-3m-4 13H7V.5H5.5"/>`),
		g.Group(children),
	)
}