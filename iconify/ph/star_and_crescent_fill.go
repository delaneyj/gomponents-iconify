package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarAndCrescentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 206.4a8 8 0 0 1-4.36 7.13A94.93 94.93 0 0 1 112 224a96 96 0 0 1 0-192a94.93 94.93 0 0 1 43.64 10.47a8 8 0 0 1 0 14.25a80 80 0 0 0 0 142.56a8 8 0 0 1 4.36 7.12Zm91.17-85.75l-26.5-11.43l-2.31-29.84a8 8 0 0 0-14.14-4.47l-18.59 22.51l-27.71-6.85a8 8 0 0 0-8.81 11.82L168.18 128l-15.07 25.61a8 8 0 0 0 8.81 11.82l27.71-6.85l18.59 22.51a8 8 0 0 0 14.14-4.47l2.31-29.84l26.5-11.43a8 8 0 0 0 0-14.7Z"/>`),
		g.Group(children),
	)
}