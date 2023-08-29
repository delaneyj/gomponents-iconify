package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M167.51 88.49a12 12 0 0 1 17-17L188 75V48a12 12 0 0 1 24 0v27l3.51-3.52a12 12 0 0 1 17 17l-24 24a12 12 0 0 1-17 0ZM244 160v40a20 20 0 0 1-20 20h-52.5a20.08 20.08 0 0 1-14.7-6.44L66.75 116H32a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h52.5a20.08 20.08 0 0 1 14.7 6.44L189.25 140H224a20 20 0 0 1 20 20Zm-24 4h-32.5a20.08 20.08 0 0 1-14.7-6.44L82.75 60H36v32h32.5a20.08 20.08 0 0 1 14.7 6.44L173.25 196H220Z"/>`),
		g.Group(children),
	)
}