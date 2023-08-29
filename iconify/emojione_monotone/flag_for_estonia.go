package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForEstonia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M6.254 21h51.492c.246.575.471 1.162.679 1.756H5.575c.208-.594.433-1.181.679-1.756M32 60C20.083 60 9.888 52.514 5.853 42h52.295C54.112 52.514 43.917 60 32 60"/>`),
		g.Group(children),
	)
}