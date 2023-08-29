package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 15H2V1h6v4h4v1h1V4L9 0H1v16h12v-2h-1v1zM9 1l3 3H9V1z"/><path fill="currentColor" d="M13 7h-2v2H9v2h2v2h2v-2h2V9h-2V7z"/>`),
		g.Group(children),
	)
}