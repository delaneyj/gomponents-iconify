package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4a1 1 0 0 1 1 1v4.2l5.213-3.65a.5.5 0 0 1 .787.41v12.08a.5.5 0 0 1-.787.41L17 14.8V19a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14ZM7.4 8.829a.4.4 0 0 0-.392.32L7 9.228v5.542a.4.4 0 0 0 .542.374l.073-.036l4.355-2.771a.401.401 0 0 0 .063-.625l-.063-.05L7.615 8.89a.4.4 0 0 0-.215-.06Z"/>`),
		g.Group(children),
	)
}