package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorUpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 42h-52.5a14 14 0 0 0-10.28 4.5L70 145.36a2 2 0 0 1-1.47.64H32a14 14 0 0 0-14 14v40a14 14 0 0 0 14 14h52.5a14 14 0 0 0 10.28-4.5L186 110.64a2 2 0 0 1 1.47-.64H224a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14Zm2 54a2 2 0 0 1-2 2h-36.5a14 14 0 0 0-10.28 4.5L86 201.36a2 2 0 0 1-1.47.64H32a2 2 0 0 1-2-2v-40a2 2 0 0 1 2-2h36.5a14 14 0 0 0 10.28-4.5L170 54.64a2 2 0 0 1 1.5-.64H224a2 2 0 0 1 2 2Zm2.24 75.76a6 6 0 1 1-8.48 8.48L206 166.49V208a6 6 0 0 1-12 0v-41.51l-13.76 13.75a6 6 0 0 1-8.48-8.48l24-24a6 6 0 0 1 8.48 0Z"/>`),
		g.Group(children),
	)
}