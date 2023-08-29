package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 36h-52.5a20.07 20.07 0 0 0-14.69 6.43L66.75 140H32a20 20 0 0 0-20 20v40a20 20 0 0 0 20 20h52.5a20.07 20.07 0 0 0 14.69-6.43L189.25 116H224a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 56h-32.5a20.07 20.07 0 0 0-14.69 6.43L82.75 196H36v-32h32.5a20.07 20.07 0 0 0 14.69-6.43L173.25 60H220Zm12.49 75.51a12 12 0 0 1-17 17L212 181v27a12 12 0 0 1-24 0v-27l-3.51 3.52a12 12 0 0 1-17-17l24-24a12 12 0 0 1 17 0Z"/>`),
		g.Group(children),
	)
}