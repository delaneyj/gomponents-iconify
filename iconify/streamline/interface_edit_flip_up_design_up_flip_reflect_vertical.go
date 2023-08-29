package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipUpDesignUpFlipReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2.5v-1a1 1 0 0 0-1-1h-1m-3 0h-3m-5 2v-1a1 1 0 0 1 1-1h1m-2 9v3a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-3m-13-4V7h13V5.5"/>`),
		g.Group(children),
	)
}