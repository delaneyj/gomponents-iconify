package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 1h16v22H3V11.586l2-2V1Zm2 2v7.414l-2 2V21h14V3H7Zm4 2v3H9V5h2Zm3 0v3h-2V5h2Zm3 0v3h-2V5h2Z"/>`),
		g.Group(children),
	)
}