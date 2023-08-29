package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomesticWorkerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M13 9.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Z"/><path fill-rule="evenodd" d="M29.449 16.05A2.006 2.006 0 0 0 29 16h-9v12h-7V16h-1a6 6 0 0 0-1 11.917V40a2 2 0 1 0 4 0v-9h3v9a2 2 0 1 0 4 0V20h7c.477 0 .914-.166 1.258-.445l1.998 8.655a4 4 0 0 0-2.998 4.797l1.05 8.994l2.824-.652l-1.03-4.914l1.958-.41l1.021 4.874l6.92-1.597l-2.999-8.544a4 4 0 0 0-4.797-2.998L30.38 11.196a1 1 0 0 0-1.95.45l1.018 4.405ZM9 22a2 2 0 0 1 2-2v4a2 2 0 0 1-2-2Zm25.655 7.709l-1.95.45a2 2 0 0 0-1.498 2.398l5.846-1.35a2 2 0 0 0-2.398-1.498Z" clip-rule="evenodd"/><path d="M18 16v3h-3v-3h3Z"/></g>`),
		g.Group(children),
	)
}