package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM5.17 4a3.001 3.001 0 0 1 5.66 0H22v2H10.83a3.001 3.001 0 0 1-5.66 0H2V4h3.17Zm8 7a3.001 3.001 0 0 1 5.66 0H22v2h-3.17a3.001 3.001 0 0 1-5.66 0H2v-2h11.17ZM16 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-8 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-2.83 0a3.001 3.001 0 0 1 5.66 0H22v2H10.83a3.001 3.001 0 0 1-5.66 0H2v-2h3.17Z"/>`),
		g.Group(children),
	)
}