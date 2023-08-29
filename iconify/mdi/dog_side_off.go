package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogSideOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 10l-3-3l4-4v2l3 3l-2 2l-1-1l-1 1m-1 1l-3-3l-1 1h-.8l4.8 4.8V11M2.39 1.73L1.11 3l6 6H5L3 7L2 8l3 3v3l-1 1v6h2v-3l2-3h5.11L15 16.89V21h2v-2.11l3.84 3.84l1.27-1.27L2.39 1.73Z"/>`),
		g.Group(children),
	)
}