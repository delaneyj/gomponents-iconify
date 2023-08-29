package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkbenchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V8.5L12 2l10 6.5V22l-10-7Zm10-9.4l6.325-4.1L12 4.4L5.675 8.5Zm-8 5.55l6.2-4.325L4 9.8Zm16 0V9.8l-6.2 4.025Zm-9.8-4.325Zm3.6 0Z"/>`),
		g.Group(children),
	)
}