package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.014 1c8.285 0 15 6.716 15 15c0 8.284-6.715 15-15 15c-8.284 0-15-6.716-15-15c0-8.284 6.716-15 15-15ZM5.491 23.634a13.05 13.05 0 0 0 5.228 4.242a3.5 3.5 0 0 0-5.228-4.242Zm18.841 2.357a13.037 13.037 0 0 0 3.671-4.955a3.5 3.5 0 0 0-3.67 4.955ZM15.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-6 9a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7ZM25 11.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0ZM19.5 21a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}