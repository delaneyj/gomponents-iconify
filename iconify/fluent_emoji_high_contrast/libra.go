package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M26 24c0 .55-.45 1-1 1H7c-.55 0-1-.45-1-1s.45-1 1-1h18c.55 0 1 .45 1 1Zm-5.1-3H25c.55 0 1-.45 1-1s-.45-1-1-1h-2.07c.7-1.2 1.07-2.57 1.07-4c0-4.41-3.59-8-8-8s-8 3.59-8 8c0 1.43.37 2.8 1.07 4H7c-.55 0-1 .45-1 1s.45 1 1 1h4.1c.41 0 .77-.25.93-.62c.15-.37.06-.81-.23-1.09A5.966 5.966 0 0 1 10 15c0-3.31 2.69-6 6-6s6 2.69 6 6c0 1.63-.64 3.15-1.8 4.29c-.29.28-.38.71-.23 1.09c.16.37.52.62.93.62Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}