package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 12a84.05 84.05 0 0 0-80.62 107.65l-55.87 55.86A12 12 0 0 0 20 184v40a12 12 0 0 0 12 12h40a12 12 0 0 0 12-12v-12h12a12 12 0 0 0 12-12v-12h12a12 12 0 0 0 8.49-3.51l7.86-7.87A84 84 0 1 0 160 12Zm0 144a59.58 59.58 0 0 1-22.1-4.2a12 12 0 0 0-13.22 2.55L115 164H96a12 12 0 0 0-12 12v12H72a12 12 0 0 0-12 12v12H44v-23l57.65-57.65a12 12 0 0 0 2.55-13.21A60 60 0 1 1 160 156Zm36-80a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}