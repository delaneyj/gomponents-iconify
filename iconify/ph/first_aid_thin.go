package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 92h-52V40a12 12 0 0 0-12-12h-48a12 12 0 0 0-12 12v52H40a12 12 0 0 0-12 12v48a12 12 0 0 0 12 12h52v52a12 12 0 0 0 12 12h48a12 12 0 0 0 12-12v-52h52a12 12 0 0 0 12-12v-48a12 12 0 0 0-12-12Zm4 60a4 4 0 0 1-4 4h-56a4 4 0 0 0-4 4v56a4 4 0 0 1-4 4h-48a4 4 0 0 1-4-4v-56a4 4 0 0 0-4-4H40a4 4 0 0 1-4-4v-48a4 4 0 0 1 4-4h56a4 4 0 0 0 4-4V40a4 4 0 0 1 4-4h48a4 4 0 0 1 4 4v56a4 4 0 0 0 4 4h56a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}