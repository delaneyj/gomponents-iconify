package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WetFloor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m17.5 1l-4.621 4.621a3 3 0 0 1-2.122.879H7.243a3 3 0 0 0-1.132.222M.5 12l4.621-4.621a3 3 0 0 1 .99-.657m5.346-.305L14 11.5h5c.936 1.123 1.591 1.965 2.596 2.528c.723.405 1.575.472 2.404.472M15.5 24c0-1.5-1-2.5-2.5-2.5m6.5 2.5c0-1.5 1-2.5 2.5-2.5m-4-1c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-.5h-3.646a3 3 0 0 1-2.683-1.658l-3.06-6.12M8.195 4.5s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 7.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61l-.291.079Z"/>`),
		g.Group(children),
	)
}