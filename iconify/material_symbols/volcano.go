package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volcano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 22l4-9h3l2-5h7l4 14H2ZM13 5V1h2v4h-2Zm4.525 1.875l-1.4-1.4L18.95 2.65l1.425 1.4l-2.85 2.825Zm-7.05 0L7.65 4.05l1.4-1.425l2.825 2.85l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}