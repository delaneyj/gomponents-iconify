package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 20.513a2.5 2.5 0 0 0 5 0V4.487a2.5 2.5 0 0 0-5 0v16.026Zm0 6.917c0 1.41 1.15 2.56 2.56 2.56c1.41 0 2.56-1.15 2.56-2.56c0-1.41-1.15-2.56-2.56-2.56c-1.41 0-2.56 1.15-2.56 2.56Zm9.56 2.56c-1.41 0-2.56-1.15-2.56-2.56c0-1.41 1.15-2.56 2.56-2.56c1.41 0 2.56 1.15 2.56 2.56c0 1.41-1.15 2.56-2.56 2.56ZM17 20.513a2.5 2.5 0 0 0 5 0V4.487a2.5 2.5 0 0 0-5 0v16.026Z"/>`),
		g.Group(children),
	)
}