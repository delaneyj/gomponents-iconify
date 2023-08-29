package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeaceBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm84 108a83.43 83.43 0 0 1-9 37.84l-63-44.09V44.87A84.12 84.12 0 0 1 212 128Zm-96-83.13v76.88l-63 44.09a83.93 83.93 0 0 1 63-121ZM66.83 185.48L116 151.05v60.08a83.86 83.86 0 0 1-49.17-25.65ZM140 211.13v-60.08l49.17 34.43A83.86 83.86 0 0 1 140 211.13Z"/>`),
		g.Group(children),
	)
}