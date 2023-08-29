package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 10a7 7 0 1 1 14 0a7 7 0 0 1-14 0Zm7-8a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm1.995 4.309c0-.738-.954-1.032-1.37-.423l-3.533 5.176a.6.6 0 0 0 .496.938h3.41v1.5a.5.5 0 0 0 1 0V12H13a.5.5 0 0 0 0-1h-1.003l-.002-4.691Zm-1 .809L10.997 11H8.345l2.65-3.882Z"/>`),
		g.Group(children),
	)
}