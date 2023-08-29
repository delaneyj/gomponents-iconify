package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M14.613 7.802a1.75 1.75 0 0 1 1.275 2.121l-1.292 5.182h1.9V14a1.75 1.75 0 1 1 3.5 0v1.14a1.75 1.75 0 0 1 0 3.43v3.81a1.75 1.75 0 1 1-3.5 0v-3.775h-4.14a1.75 1.75 0 0 1-1.698-2.173l1.834-7.355a1.75 1.75 0 0 1 2.121-1.275Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}