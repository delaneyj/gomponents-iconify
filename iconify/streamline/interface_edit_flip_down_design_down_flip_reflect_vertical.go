package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipDownDesignDownFlipReflectVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5v1a1 1 0 0 1-1 1h-1m-3 0h-3m-5-2v1a1 1 0 0 0 1 1h1m-2-9v-3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v3m-13 4V7h13v1.5"/>`),
		g.Group(children),
	)
}