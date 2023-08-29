package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteeringWheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipOSteeringWheel0" d="M32 24a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z"/></defs><g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm0 0V32M4 24h12m28 0H32"/><use href="#ipOSteeringWheel0"/><use href="#ipOSteeringWheel0" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 24a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z"/></g>`),
		g.Group(children),
	)
}