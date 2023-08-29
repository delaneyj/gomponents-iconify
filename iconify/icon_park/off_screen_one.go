package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OffScreenOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 6L16 15.8995"/><path d="M6 41.8995L16 32"/><path d="M42.0001 41.8995L32.1006 32"/><path d="M41.8995 6L32 15.8995"/><path d="M32 7V16H41"/><path d="M16 7V16H7"/><path d="M16 41V32H7"/><path d="M32 41V32H40.8995"/></g>`),
		g.Group(children),
	)
}