package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.31 6.51c.843.728.925 1.988.182 2.814a2.05 2.05 0 0 1-1.526.676H2.034C.911 10 0 9.107 0 8.006c0-.573.251-1.118.69-1.496L7.656.498a2.065 2.065 0 0 1 2.688 0L17.31 6.51zM2.034 8.006h13.932L9 1.994L2.034 8.006zM18 14v2a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2zm-2 0H2v2h14v-2z"/>`),
		g.Group(children),
	)
}