package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandscapeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18q-.625 0-.9-.55t.1-1.05l4-5.325q.3-.4.8-.4t.8.4l3.1 4.125q.25.35.65.4t.75-.2q.325-.25.388-.625t-.138-.725L10.25 11l2.95-3.925q.3-.4.8-.4t.8.4l7 9.325q.375.5.1 1.05T21 18H3Z"/>`),
		g.Group(children),
	)
}