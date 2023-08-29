package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M14 8a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v9.75c0 .138.112.25.25.25h3.063c.9 0 1.342 1.096.693 1.72l-5.313 5.113a1 1 0 0 1-1.386 0L9.994 19.72c-.649-.625-.207-1.721.693-1.721h3.063a.25.25 0 0 0 .25-.25V8Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}