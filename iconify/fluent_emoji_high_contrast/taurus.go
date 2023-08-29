package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taurus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17.99 8.75C19.63 7.3 21.76 6.5 24 6.5c.55 0 1 .45 1 1s-.45 1-1 1c-1.74 0-3.4.61-4.67 1.73A8.01 8.01 0 0 1 24 17.5c0 4.41-3.59 8-8 8s-8-3.59-8-8a8.01 8.01 0 0 1 4.67-7.27C11.4 9.11 9.74 8.5 8 8.5c-.55 0-1-.45-1-1s.45-1 1-1c2.24 0 4.37.8 6.01 2.25c.55.48 1.26.75 1.99.75s1.44-.27 1.99-.75ZM16 23.5c3.31 0 6-2.69 6-6s-2.69-6-6-6s-6 2.69-6 6s2.69 6 6 6Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}