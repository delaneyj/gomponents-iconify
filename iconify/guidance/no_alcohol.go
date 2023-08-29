package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoAlcohol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.5 16v7.5H20m1.5-2v-8l-1.8-2.4a6 6 0 0 1-1.2-3.6v-7h-3v7a6 6 0 0 1-1.2 3.6L13 12.833l-.083.084M5 19.5a3.5 3.5 0 0 1-3.394-4.364L2.5 11.5h5l.894 3.636A3.5 3.5 0 0 1 5 19.5Zm0 0v4m0 0h3m-3 0H2M.5.5l12.417 12.417M23.5 23.5L12.917 12.917"/>`),
		g.Group(children),
	)
}