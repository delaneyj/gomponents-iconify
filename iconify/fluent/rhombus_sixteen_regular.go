package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.892 4.069A1.75 1.75 0 0 1 5.504 3h7.748a1.75 1.75 0 0 1 1.611 2.431l-2.747 6.502a1.75 1.75 0 0 1-1.612 1.069H2.756a1.75 1.75 0 0 1-1.612-2.432l2.748-6.5ZM5.504 4a.75.75 0 0 0-.691.458L2.065 10.96a.75.75 0 0 0 .69 1.042h7.749a.75.75 0 0 0 .69-.458l2.748-6.502A.75.75 0 0 0 13.252 4H5.503Z"/>`),
		g.Group(children),
	)
}