package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h16v7.414l-3 3V24H7V10.414l-3-3V0Zm2 2v1.5h12V2H6Zm12 3.5H6v1.086l3 3V22h6V9.586l3-3V5.5ZM11 10h2.004v2.004H11V10Z"/>`),
		g.Group(children),
	)
}