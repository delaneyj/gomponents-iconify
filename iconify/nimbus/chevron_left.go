package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m5.82 8l5.66-5.56l-.87-.89L4.9 7.09a1.18 1.18 0 0 0-.39.91a1.13 1.13 0 0 0 .39.85l5.7 5.7l.89-.88z"/>`),
		g.Group(children),
	)
}