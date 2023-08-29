package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h14.414L22 7.586V22H2V2Zm2 2v16h2v-6h12v6h2V8.414L15.586 4H13v4H6V4H4Zm4 0v2h3V4H8Zm8 16v-4H8v4h8Z"/>`),
		g.Group(children),
	)
}