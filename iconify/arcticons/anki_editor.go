package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnkiEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.3 18.1l3.7 6.2l7.3-.4l-4.8 5.4l2.9 6.8l-6.6-3.1l-5.6 4.9l.6-7.3l-6.1-3.7l7.1-1.6l1.5-7.2zm14.1-8.4l-.4 3.5l2.7 1.7l-3.1.5l-1 3.2l-1.6-2.8l-3.3.1l2.2-2.4l-1.1-3.1l3.1 1.5l2.5-2.2z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}