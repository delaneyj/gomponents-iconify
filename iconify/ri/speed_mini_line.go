package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeedMiniLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.032 12L6 9.86v4.28L9.032 12Zm-4.244 5.443A.5.5 0 0 1 4 17.035V6.965a.5.5 0 0 1 .788-.409l7.133 5.035a.5.5 0 0 1 0 .817l-7.133 5.035ZM15 14.14L18.032 12L15 9.86v4.28Zm-2-7.175a.5.5 0 0 1 .788-.409l7.133 5.035a.5.5 0 0 1 0 .817l-7.133 5.035a.5.5 0 0 1-.788-.408V6.965Z"/>`),
		g.Group(children),
	)
}