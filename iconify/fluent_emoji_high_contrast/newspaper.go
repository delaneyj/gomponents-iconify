package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M6 10a1 1 0 0 0 0 2h18a1 1 0 1 0 0-2H6Zm-.5 4a.5.5 0 0 0 0 1h19a.5.5 0 1 0 0-1h-19ZM19 17.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h5a.5.5 0 1 0 0-1h-5Zm-.5 3.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h5a.5.5 0 1 0 0-1h-5ZM5 19a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-6Z"/><path d="M8 4a2 2 0 0 0-2 2H3.962C2.252 6 1 7.418 1 9v18c0 2.134 1.683 4 3.923 4h24.094A2 2 0 0 0 31 29V6a2 2 0 0 0-2-2H8Zm18.038 4c.46 0 .846.336.94.786c.014.069.022.14.022.214v18a2 2 0 0 0 2 2H4.923C3.861 29 3 28.105 3 27V9c0-.552.43-1 .962-1h22.076Z"/></g>`),
		g.Group(children),
	)
}