package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M7.417 16.944a1 1 0 0 1 0-1.387l5.113-5.313c.624-.649 1.72-.207 1.72.693V14c0 .138.112.25.25.25h9.75a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H14.5a.25.25 0 0 0-.25.25v3.063c0 .9-1.096 1.342-1.72.693l-5.113-5.313Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}