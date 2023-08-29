package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M20.01 6.52c.76-.66 1.74-1.02 2.75-1.02C25.1 5.5 27 7.4 27 9.74c0 .55-.45 1-1 1s-1-.45-1-1c0-1.24-1-2.24-2.24-2.24c-1.11 0-2.04.8-2.21 1.89l-2.41 15.28A2.154 2.154 0 0 1 16 26.5c-1.07 0-1.97-.77-2.14-1.83L11.45 9.39A2.22 2.22 0 0 0 9.24 7.5C8 7.5 7 8.5 7 9.74c0 .55-.45 1-1 1s-1-.45-1-1a4.242 4.242 0 0 1 6.99-3.23c.76.66 1.27 1.56 1.43 2.56l2.41 15.29c.01.08.09.14.17.14c.09 0 .16-.06.17-.14l2.41-15.28c.16-1 .66-1.91 1.43-2.56Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}