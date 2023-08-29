package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.703 1h-2.828c-.223 0-.553.105-.69.381s-.31.619-.31.619l-5 10l3 6l.31.619c.138.275.467.381.69.381h2.828c.5 0 .48-.381.355-.631L15.875 18l-3-6l5-10l.184-.369c.125-.25.144-.631-.356-.631zM6.815 5.381C6.678 5.105 6.348 5 6.125 5H3.297c-.5 0-.48.381-.355.631L3.125 6l1.25 2.5l-2.25 4.5l-.184.369c-.125.25-.144.631.356.631h2.828c.223 0 .553-.106.691-.381L6.125 13l2.25-4.5L7.125 6l-.31-.619z"/>`),
		g.Group(children),
	)
}