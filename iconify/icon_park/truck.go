package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M12 39C14.2091 39 16 37.2091 16 35C16 32.7909 14.2091 31 12 31C9.79086 31 8 32.7909 8 35C8 37.2091 9.79086 39 12 39Z"/><path d="M35 39C37.2091 39 39 37.2091 39 35C39 32.7909 37.2091 31 35 31C32.7909 31 31 32.7909 31 35C31 37.2091 32.7909 39 35 39Z"/><path stroke-linecap="round" d="M8 35H2V11H31V35H16"/><path stroke-linecap="round" d="M31 35V18H39.5714L46 26.5V35H39.8112"/></g>`),
		g.Group(children),
	)
}