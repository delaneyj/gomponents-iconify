package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-32 40h-16.08a16 16 0 0 0-15.73 13l-6.55 35H168a8 8 0 0 1 0 16h-33.36l-7.11 37.9A32 32 0 0 1 96.08 200H80a8 8 0 0 1 0-16h16.08a16 16 0 0 0 15.73-13l6.55-35H88a8 8 0 0 1 0-16h33.36l7.11-37.9A32 32 0 0 1 159.92 56H176a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}