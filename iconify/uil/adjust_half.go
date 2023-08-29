package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.83 7.32a10.11 10.11 0 0 0-3.44-3.73A10 10 0 1 0 12 22h.29a10 10 0 0 0 8.54-14.68ZM11 19.93a8 8 0 0 1 0-15.86Zm2-15.86a8.07 8.07 0 0 1 2.49.74L13 9.12Zm0 9l4.17-7.17a8.14 8.14 0 0 1 1.58 1.83L13 17.69Zm1.15 6.58L19.74 10a8.16 8.16 0 0 1 .26 2a8 8 0 0 1-5.85 7.7Z"/>`),
		g.Group(children),
	)
}