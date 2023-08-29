package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTransferDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 464h480v32H16zm379.313-164.687l-22.626-22.626L280 369.373V16h-32v353.373l-92.687-92.686l-22.626 22.626L264 430.627l131.313-131.314z"/>`),
		g.Group(children),
	)
}