package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saopaulometro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.366 6.543l5.092 5.456l-5.092 5.456V6.543zM24 0v24H0V0h24zm-5.542 11.999l1.747-1.872L11.976 1.9l-8.227 8.228l1.747 1.871l-1.747 1.871l8.227 8.229l8.228-8.229l-1.746-1.871zm-7.87 5.455V6.543l-5.092 5.456l5.092 5.455z"/>`),
		g.Group(children),
	)
}