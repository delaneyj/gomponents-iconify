package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 6.6L8 1.4L6 2.7V1H4v3L0 6.6l1.9 2.7l.1-.1V15h5v-4h2v4h5V9.2l.1.1L16 6.6zm-14.6.3L8 2.6l6.6 4.3l-.7 1L8 4L2.1 7.9l-.7-1zM13 14h-3v-4H6v4H3V8.6l5-3.3l5 3.3V14z"/>`),
		g.Group(children),
	)
}