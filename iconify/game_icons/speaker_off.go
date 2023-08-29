package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m275.5 96l-96 96h-96v128h96l96 96V96zm50.863 89.637l-12.726 12.726L371.273 256l-57.636 57.637l12.726 12.726L384 268.727l57.637 57.636l12.726-12.726L396.727 256l57.636-57.637l-12.726-12.726L384 243.273l-57.637-57.636z"/>`),
		g.Group(children),
	)
}