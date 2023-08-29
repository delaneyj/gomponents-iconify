package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackwardAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.29 4.31a.776.776 0 0 0-.82.16l-6.72 6.72V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v14c0 .41.34.75.75.75s.75-.34.75-.75v-6.19l6.72 6.72c.14.14.34.22.53.22a.753.753 0 0 0 .75-.75V5c0-.3-.18-.58-.46-.69Zm-1.04 12.88L10.06 12l5.19-5.19v10.38Z"/>`),
		g.Group(children),
	)
}