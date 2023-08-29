package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareZeroBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 188c28.26 0 48-24.67 48-60s-19.74-60-48-60s-48 24.67-48 60s19.74 60 48 60Zm0-96c23.33 0 24 32.32 24 36s-.67 36-24 36s-24-32.32-24-36s.67-36 24-36Zm80-64H48a20 20 0 0 0-20 20v160a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-4 176H52V52h152Z"/>`),
		g.Group(children),
	)
}