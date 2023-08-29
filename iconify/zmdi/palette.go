package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 0q80 0 136 50t56 121q0 44-31.5 75T277 277h-37q-14 0-23 9.5t-9 22.5q0 12 8 21q8 10 8 22q0 13-9.5 22.5T192 384q-80 0-136-56T0 192T56 56T192 0zM74.5 192q13.5 0 23-9.5T107 160t-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5t9 22.5t22.5 9.5zm64-85q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zm107 0q13.5 0 22.5-9.5t9-23t-9-22.5t-22.5-9t-23 9t-9.5 22.5t9.5 23t23 9.5zm64 85q13.5 0 22.5-9.5t9-22.5t-9-22.5t-22.5-9.5t-23 9.5T277 160t9.5 22.5t23 9.5z"/>`),
		g.Group(children),
	)
}