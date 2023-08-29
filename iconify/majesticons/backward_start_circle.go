package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackwardStartCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1zm5.493 7.13a1 1 0 0 0-1.008.013L13 10.233V9a1 1 0 0 0-1.514-.857L8 10.233V9a1 1 0 0 0-2 0v6a1 1 0 1 0 2 0v-1.234l3.486 2.091A1 1 0 0 0 13 15v-1.234l3.485 2.091A1 1 0 0 0 18 15V9a1 1 0 0 0-.507-.87zM11 12v-1.234L8.944 12L11 13.234V12zm5 1.234L13.944 12L16 10.766v2.468z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}