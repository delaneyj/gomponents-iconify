package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeVscodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#007acc" d="m11.72 18.685l-6.883 5.184L2 22.922L9 16L2 9.077l2.837-.947l6.883 5.184l11.433-11.3L30 4.922v22.155l-6.847 2.909ZM15.287 16l7.865 5.923V10.076Z"/>`),
		g.Group(children),
	)
}