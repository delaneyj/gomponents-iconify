package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm0 176H48V48h160v160ZM168 96v64a8 8 0 0 1-8 8H96a8 8 0 0 1-8-8V96a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}