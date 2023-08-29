package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaySkipBackSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M143.47 64v163.52L416 64v384L143.47 284.48V448H96V64h47.47z"/>`),
		g.Group(children),
	)
}