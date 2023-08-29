package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachmentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 0v4.5a2 2 0 1 0 4 0v-3a1 1 0 0 0-2 0V5M6 .5h6.5a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1V8M11 4.5H7m4 3H7m4 3H4"/>`),
		g.Group(children),
	)
}