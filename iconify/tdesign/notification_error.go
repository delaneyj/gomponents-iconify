package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8a8 8 0 1 1 16 0v4.697l2 3V20h-5.611a4.502 4.502 0 0 1-8.777 0H2v-4.303l2-3V8Zm16 10v-1.697l-2-3V8A6 6 0 0 0 6 8v5.303l-2 3V18h16Zm-5.708 2H9.708a2.5 2.5 0 0 0 4.584 0ZM13 6v6h-2V6h2Zm-2 7.5h2.004v2.004H11V13.5Z"/>`),
		g.Group(children),
	)
}