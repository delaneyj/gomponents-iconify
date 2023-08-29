package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16.5 18.5c0 .818-.393 1.544-1 2V24a1.5 1.5 0 0 1-3 0v-3.5a2.5 2.5 0 1 1 4-2Z"/><path d="M17.5 7a6 6 0 0 1 12 0v1.75a1.25 1.25 0 1 1-2.5 0V7a3.5 3.5 0 1 0-7 0v3h2a4 4 0 0 1 4 4v13a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V14a4 4 0 0 1 4-4h11.5V7ZM6 12a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V14a2 2 0 0 0-2-2H6Z"/></g>`),
		g.Group(children),
	)
}