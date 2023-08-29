package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextMoreTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.672 2.611a1 1 0 0 0-1.843 0l-6.75 16a1 1 0 0 0 1.843.778L5.773 15h7.954l1.286 3.047A2.496 2.496 0 0 1 17.5 19a.997.997 0 0 0-.078-.389l-6.75-16ZM12.884 13H6.617L9.75 5.573L12.884 13ZM10.5 22a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm6.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}