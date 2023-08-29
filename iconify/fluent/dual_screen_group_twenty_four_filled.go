package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenGroupTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.253 4.004c.966 0 1.75.784 1.75 1.75V18.25a1.75 1.75 0 0 1-1.75 1.75h-7.248a1.79 1.79 0 0 1-.255-.018V4.023a1.79 1.79 0 0 1 .255-.019h7.248Zm-9.248-.002c.084 0 .166.006.246.017V19.98c-.08.011-.162.017-.246.017H3.758a1.75 1.75 0 0 1-1.75-1.75V5.752c0-.967.783-1.75 1.75-1.75h7.247Z"/>`),
		g.Group(children),
	)
}