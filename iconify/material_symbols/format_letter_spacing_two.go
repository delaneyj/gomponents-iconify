package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLetterSpacingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 22l-4-4l4-4l1.425 1.4l-1.6 1.6h12.35L16.6 15.4L18 14l4 4l-4 4l-1.425-1.4l1.6-1.6H5.825L7.4 20.6L6 22Zm.9-9L11 2h2l4.1 11h-1.9l-.95-2.8H9.8l-1 2.8H6.9Zm3.45-4.4h3.3l-1.6-4.55h-.1l-1.6 4.55Z"/>`),
		g.Group(children),
	)
}