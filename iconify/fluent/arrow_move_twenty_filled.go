package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMoveTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.53 2.22a.75.75 0 0 0-1.06 0L7.22 4.47a.75.75 0 0 0 1.06 1.06l.97-.97v2.69a.75.75 0 0 0 1.5 0V4.56l.97.97a.75.75 0 1 0 1.06-1.06l-2.25-2.25ZM2.22 9.47a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 0 0 1.06-1.06l-.97-.97h2.69a.75.75 0 0 0 0-1.5H4.56l.97-.97a.75.75 0 0 0-1.06-1.06L2.22 9.47Zm8.31 8.31l2.25-2.25a.75.75 0 1 0-1.06-1.06l-.97.97v-2.69a.75.75 0 0 0-1.5 0v2.69l-.97-.97a.75.75 0 0 0-1.06 1.06l2.25 2.25a.75.75 0 0 0 1.06 0Zm5-10.56l2.25 2.25a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 1 1-1.06-1.06l.97-.97h-2.69a.75.75 0 0 1 0-1.5h2.69l-.97-.97a.75.75 0 0 1 1.06-1.06Z"/>`),
		g.Group(children),
	)
}