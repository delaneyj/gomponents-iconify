package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaySkipForwardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M368.53 64v163.52L96 64v384l272.53-163.52V448H416V64h-47.47z"/>`),
		g.Group(children),
	)
}