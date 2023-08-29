package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 15H2V1h6v4h4v2.59l1-1V4L9 0H1v16h12v-2.59l-1-1V15zM9 1l3 3H9V1z"/><path fill="currentColor" d="m15 8l-1-1l-2 2l-2-2l-1 1l2 2l-2 2l1 1l2-2l2 2l1-1l-2-2l2-2z"/>`),
		g.Group(children),
	)
}