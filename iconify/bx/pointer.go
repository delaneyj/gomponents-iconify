package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pointer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.978 13.21a1 1 0 0 0-.396-1.024l-14-10a.999.999 0 0 0-1.575.931l2 17a1 1 0 0 0 1.767.516l3.612-4.416l3.377 5.46l1.701-1.052l-3.357-5.428l6.089-1.218a.995.995 0 0 0 .782-.769zm-8.674.31a1 1 0 0 0-.578.347l-3.008 3.677L7.257 5.127l10.283 7.345l-5.236 1.048z"/>`),
		g.Group(children),
	)
}