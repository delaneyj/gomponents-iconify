package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontFamily(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.554 22H3.4L11 3h2l7.6 19h-2.154l-2.4-6H7.954l-2.4 6Zm3.2-8h6.492L12 5.885L8.754 14Z"/>`),
		g.Group(children),
	)
}