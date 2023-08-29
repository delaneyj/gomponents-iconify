package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareFourBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 28H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152ZM86.2 150.93A12 12 0 0 1 84.68 140l24-68a12 12 0 1 1 22.64 8L113 132h19v-16a12 12 0 0 1 24 0v16h4a12 12 0 0 1 0 24h-4v20a12 12 0 0 1-24 0v-20H96a12 12 0 0 1-9.8-5.07Z"/>`),
		g.Group(children),
	)
}