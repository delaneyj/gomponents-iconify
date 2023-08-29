package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screencast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2h22v17h-6v-2h4V4H3l.001 13h4v2H1V2Zm4.586 20L12 15.585L18.414 22H5.586Zm4.828-2h3.172L12 18.414L10.414 20Z"/>`),
		g.Group(children),
	)
}