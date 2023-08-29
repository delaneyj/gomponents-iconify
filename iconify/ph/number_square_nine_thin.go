package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareNineThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a12 12 0 0 0-12 12v160a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12Zm4 172a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V48a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4ZM144 80.29a32 32 0 1 0-6.23 58.16L116.57 174a4 4 0 0 0 1.38 5.48a3.92 3.92 0 0 0 2 .57a4 4 0 0 0 3.43-1.95l32.33-54.1A32 32 0 0 0 144 80.29Zm4.8 39.71l-.06.09a22.62 22.62 0 1 1 .06-.09Z"/>`),
		g.Group(children),
	)
}