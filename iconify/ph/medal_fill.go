package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 96a88 88 0 1 0-144 67.83V240a8 8 0 0 0 11.58 7.16L128 225l44.43 22.21a8.07 8.07 0 0 0 3.57.79a8 8 0 0 0 8-8v-76.17A87.85 87.85 0 0 0 216 96ZM56 96a72 72 0 1 1 72 72a72.08 72.08 0 0 1-72-72Zm16 0a56 56 0 1 1 56 56a56.06 56.06 0 0 1-56-56Z"/>`),
		g.Group(children),
	)
}