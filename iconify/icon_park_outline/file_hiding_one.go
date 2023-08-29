package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHidingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12"/><path d="M25 30a8.64 8.64 0 0 0 1.255 2.517C27.783 34.63 30.235 36 33 36s5.217-1.37 6.745-3.483A8.64 8.64 0 0 0 41 30m-10.479 6.068l-1.035 3.864m6-3.864l1.036 3.864m2.831-5.578l2.829 2.828M24 37.01l2.828-2.828M30 4v10h10"/></g>`),
		g.Group(children),
	)
}