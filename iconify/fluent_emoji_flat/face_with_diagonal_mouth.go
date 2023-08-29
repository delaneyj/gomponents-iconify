package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithDiagonalMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FFB02E" d="M15.84 29.999c9.334 0 13.998-6.268 13.998-13.999c0-7.731-4.664-13.999-13.999-13.999C6.505 2.001 1.84 8.27 1.84 16c0 7.731 4.665 13.999 14 13.999Z"/><path fill="#fff" d="M14.84 11.5a4.5 4.5 0 1 1-9.002 0a4.5 4.5 0 0 1 9.001 0Zm1.998 0a4.5 4.5 0 1 0 9.001 0a4.5 4.5 0 0 0-9 0Z"/><path fill="#402A32" d="M10.96 14.69a3.21 3.21 0 1 0 0-6.42a3.21 3.21 0 0 0 0 6.42Zm9.757 0a3.21 3.21 0 1 1 0-6.42a3.21 3.21 0 0 1 0 6.42Zm-9.621 5.591a.75.75 0 1 0 .426 1.438l9.06-2.685a.75.75 0 1 0-.426-1.438l-9.06 2.685Z"/></g>`),
		g.Group(children),
	)
}