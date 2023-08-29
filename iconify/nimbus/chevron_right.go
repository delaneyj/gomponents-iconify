package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m10.18 8.05l-5.66 5.56l.87.89l5.71-5.59a1.18 1.18 0 0 0 .39-.86a1.13 1.13 0 0 0-.39-.85L5.4 1.5l-.89.88z"/>`),
		g.Group(children),
	)
}