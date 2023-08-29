package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.4 13.8a2.2 2.2 0 1 1-4.4 0a2.2 2.2 0 0 1 4.4 0z"/><path fill="currentColor" d="M10.6 16H7.5c0-4.1-3.4-7.5-7.5-7.5V5.4c5.9 0 10.6 4.7 10.6 10.6z"/><path fill="currentColor" d="M12.8 16C12.8 8.9 7.1 3.2 0 3.2V0c8.8 0 16 7.2 16 16h-3.2z"/>`),
		g.Group(children),
	)
}