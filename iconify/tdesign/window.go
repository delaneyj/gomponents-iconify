package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h7V4H4Zm9 1.414v6.172l7 7v-6.172l-7-7Zm7 4.172V4h-5.586L20 9.586ZM18.586 20L13 14.414V20h5.586Z"/>`),
		g.Group(children),
	)
}