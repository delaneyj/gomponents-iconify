package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLetterSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 22l-4-4l4-4l1.425 1.4l-1.6 1.6h8.35L14.6 15.4L16 14l4 4l-4 4l-1.425-1.4l1.6-1.6h-8.35L9.4 20.6L8 22ZM5 12V2h2v10H5Zm6 0V2h2v10h-2Zm6 0V2h2v10h-2Z"/>`),
		g.Group(children),
	)
}