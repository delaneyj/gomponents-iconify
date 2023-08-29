package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FixedWidthTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 4c.414 0 .75.308.75.688V6h6.75V4.687c0-.38.336-.687.75-.687s.75.308.75.688V6h6.75V4.687c0-.38.336-.687.75-.687s.75.308.75.688v4.125c0 .38-.336.687-.75.687s-.75-.308-.75-.688V7.5h-6.75v1.313c0 .38-.336.687-.75.687s-.75-.308-.75-.688V7.5H4.5v1.313c0 .38-.336.687-.75.687S3 9.192 3 8.812V4.689c0-.38.336-.688.75-.688Zm2 7A2.75 2.75 0 0 0 3 13.75v4a2.75 2.75 0 0 0 2.75 2.75h5.5V11h-5.5Zm12.5 9.5h-5.5V11h5.5A2.75 2.75 0 0 1 21 13.75v4a2.75 2.75 0 0 1-2.75 2.75Z"/>`),
		g.Group(children),
	)
}