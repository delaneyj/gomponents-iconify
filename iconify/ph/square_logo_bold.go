package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM96 172h64a12 12 0 0 0 12-12V96a12 12 0 0 0-12-12H96a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12Zm12-64h40v40h-40Z"/>`),
		g.Group(children),
	)
}