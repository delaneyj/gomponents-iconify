package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.375 0h8.678l8.735 15.165H18.16zm1.446 16.393h16.607L23.089 24H6.482zM8.25 1.874l4.446 7.607L4.339 24L0 16.392z"/>`),
		g.Group(children),
	)
}