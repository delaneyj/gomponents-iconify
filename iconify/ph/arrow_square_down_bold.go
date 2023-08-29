package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM87.51 144.49a12 12 0 1 1 17-17L116 139V88a12 12 0 0 1 24 0v51l11.51-11.52a12 12 0 1 1 17 17l-32 32a12 12 0 0 1-17 0Z"/>`),
		g.Group(children),
	)
}