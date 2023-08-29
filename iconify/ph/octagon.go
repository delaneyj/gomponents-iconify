package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m227.31 80.24l-51.55-51.55A15.86 15.86 0 0 0 164.45 24h-72.9a15.86 15.86 0 0 0-11.31 4.69L28.69 80.24A15.86 15.86 0 0 0 24 91.55v72.9a15.86 15.86 0 0 0 4.69 11.31l51.55 51.55A15.86 15.86 0 0 0 91.55 232h72.9a15.86 15.86 0 0 0 11.31-4.69l51.55-51.55a15.86 15.86 0 0 0 4.69-11.31v-72.9a15.86 15.86 0 0 0-4.69-11.31ZM216 164.45L164.45 216h-72.9L40 164.45v-72.9L91.55 40h72.9L216 91.55Z"/>`),
		g.Group(children),
	)
}