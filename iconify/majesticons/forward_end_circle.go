package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardEndCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1zM6.507 8.13a1 1 0 0 1 1.008.013L11 10.233V9a1 1 0 0 1 1.514-.857L16 10.233V9a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0v-1.234l-3.486 2.091A1 1 0 0 1 11 15v-1.234l-3.486 2.091A1 1 0 0 1 6 15V9a1 1 0 0 1 .507-.87zM13 12v1.234L15.056 12L13 10.766V12zm-5 1.234L10.056 12L8 10.766v2.468z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}