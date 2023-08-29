package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiplyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M513 76.8L436.2 0L257 179.2L77.8 0L1 76.8L180.2 256L1 435.2L77.8 512L257 332.8L436.2 512l76.8-76.8L333.8 256z"/>`),
		g.Group(children),
	)
}