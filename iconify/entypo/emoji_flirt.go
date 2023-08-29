package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiFlirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 9.75C8.329 9.75 9 8.967 9 8s-.671-1.75-1.5-1.75S6 7.034 6 8s.672 1.75 1.5 1.75zM10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 0 1 0 15.2zm4.341-6.263a.756.756 0 0 0-1.008.32c-.034.065-.869 1.593-3.332 1.593c-2.451 0-3.291-1.513-3.333-1.592a.75.75 0 0 0-1.339.678c.05.099 1.248 2.414 4.672 2.414c3.425 0 4.621-2.316 4.67-2.415a.745.745 0 0 0-.33-.998zM11.25 8.75h2.5a.75.75 0 0 0 0-1.5h-2.5a.75.75 0 0 0 0 1.5z"/>`),
		g.Group(children),
	)
}