package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#9B9B9B" d="M16.055 1.5c-8.008 0-14.5 6.492-14.5 14.5s6.492 14.5 14.5 14.5c8.009 0 14.5-6.492 14.5-14.5s-6.491-14.5-14.5-14.5Z"/><path fill="#fff" d="M2.556 16c0-7.456 6.044-13.5 13.5-13.5c7.455 0 13.5 6.044 13.5 13.5s-6.045 13.5-13.5 13.5c-7.456 0-13.5-6.044-13.5-13.5Z"/></g>`),
		g.Group(children),
	)
}