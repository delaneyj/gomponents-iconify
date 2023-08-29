package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePptTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-4V3ZM2.859 2.878l12.57-1.796a.5.5 0 0 1 .571.495v20.847a.5.5 0 0 1-.57.495L2.858 21.123a1 1 0 0 1-.859-.99V3.868a1 1 0 0 1 .859-.99ZM5 8v8h2v-2h6V8H5Zm2 2h4v2H7v-2Z"/>`),
		g.Group(children),
	)
}