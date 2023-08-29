package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="36" height="38" x="6" y="5" stroke-linejoin="round"/><rect width="24" height="6" x="12" y="12" stroke-linejoin="round"/><path d="M16 18V43"/><path d="M32 18V43"/><path d="M16 24H32"/><path d="M16 30H32"/><path d="M16 36H32"/></g>`),
		g.Group(children),
	)
}