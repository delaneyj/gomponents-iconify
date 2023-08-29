package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 76H32a12 12 0 0 0-12 12v48a100.24 100.24 0 0 0 26.73 68H32a12 12 0 0 0 0 24h176a12 12 0 0 0 0-24h-14.73a100.75 100.75 0 0 0 20-32A44 44 0 0 0 256 128v-8a44.05 44.05 0 0 0-44-44Zm-16 60a76.27 76.27 0 0 1-42 68H86a76.27 76.27 0 0 1-42-68v-36h152Zm36-8a20 20 0 0 1-12.57 18.55A97.17 97.17 0 0 0 220 136v-34.32A20 20 0 0 1 232 120ZM68 48V24a12 12 0 0 1 24 0v24a12 12 0 0 1-24 0Zm40 0V24a12 12 0 0 1 24 0v24a12 12 0 0 1-24 0Zm40 0V24a12 12 0 0 1 24 0v24a12 12 0 0 1-24 0Z"/>`),
		g.Group(children),
	)
}