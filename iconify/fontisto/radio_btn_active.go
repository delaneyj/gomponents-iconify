package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioBtnActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 12C0 5.373 5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12C5.376 23.992.008 18.624 0 12.001zm2.4 0c0 5.302 4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6s-4.298-9.6-9.6-9.6c-5.299.006-9.594 4.301-9.6 9.599V12zm4 0a5.6 5.6 0 1 1 11.2 0a5.6 5.6 0 0 1-11.2 0z"/>`),
		g.Group(children),
	)
}