package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarricadeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 66H32a14 14 0 0 0-14 14v72a14 14 0 0 0 14 14h26v34a6 6 0 0 0 12 0v-34h116v34a6 6 0 0 0 12 0v-34h26a14 14 0 0 0 14-14V80a14 14 0 0 0-14-14Zm2 14v53.52L170.48 78H224a2 2 0 0 1 2 2ZM32 78h49.52l76 76h-55L30 81.52V80a2 2 0 0 1 2-2Zm-2 74V98.48L85.52 154H32a2 2 0 0 1-2-2Zm194 2h-49.52l-76-76h55L226 150.48V152a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}