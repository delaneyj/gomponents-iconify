package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LotionBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17 14s-1-2-2 0c0 0-3 4.5-3 7.5c0 1 .5 3.5 4 3.5s4-2.5 4-3.5c0-3-3-7.5-3-7.5Z"/><path d="M15 2c-2 0-3 2-1 3v1.5c0 .466-.438.621-.498.64A8.003 8.003 0 0 0 7 15v11a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4V15c0-3.906-2.8-7.159-6.502-7.86h.006S18.004 7 18 6.5V5h3c1 0 1.5-1.5 0-2c0 0-5-1-6-1Zm0 7h2a6 6 0 0 1 6 6v11a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V15a6 6 0 0 1 6-6Z"/></g>`),
		g.Group(children),
	)
}