package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2V2Zm13 0h7v7h-2V4h-5V2ZM2 11h20v2H2v-2Zm2 4v5h5v2H2v-7h2Zm18 0v7h-7v-2h5v-5h2Z"/>`),
		g.Group(children),
	)
}