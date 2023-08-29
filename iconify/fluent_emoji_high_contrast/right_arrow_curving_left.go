package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M26 14.5c0-4.141-3.132-7.498-6.997-7.5H13a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h6c1.401 0 3 1.303 3 3.5S20.401 18 19 18h-6.75a.25.25 0 0 1-.25-.25v-2.246a.75.75 0 0 0-1.3-.509l-4.23 4.578a.75.75 0 0 0 0 1.018l4.23 4.578a.75.75 0 0 0 1.3-.509v-2.41a.25.25 0 0 1 .25-.25H19c3.866 0 7-3.358 7-7.5Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}