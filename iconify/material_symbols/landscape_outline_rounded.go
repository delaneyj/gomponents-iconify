package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandscapeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18q-.625 0-.9-.55t.1-1.05l4-5.325q.3-.4.8-.4t.8.4L11.5 16H19l-5-6.65l-2.5 3.3L10.25 11l2.95-3.925q.3-.4.8-.4t.8.4l7 9.325q.375.5.1 1.05T21 18H3Zm11.025-2ZM5 16h4l-2-2.675L5 16Zm0 0h4h-4Z"/>`),
		g.Group(children),
	)
}