package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.801 23.999h-3.6a1.202 1.202 0 0 1-1.2-1.2v-21.6c0-.663.537-1.2 1.2-1.201h3.6a1.202 1.202 0 0 1 1.2 1.2v21.6c0 .663-.537 1.2-1.2 1.201zm11.999 0h-3.6a1.2 1.2 0 0 1-1.2-1.2V1.198a1.2 1.2 0 0 1 1.2-1.2h3.6a1.2 1.2 0 0 1 1.2 1.2v21.6a1.2 1.2 0 0 1-1.2 1.2z"/>`),
		g.Group(children),
	)
}