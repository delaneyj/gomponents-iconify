package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 12.25A6.25 6.25 0 0 1 12.25 6h23.5A6.25 6.25 0 0 1 42 12.25V14H6v-1.75Zm0 4.25v19.25A6.25 6.25 0 0 0 12.25 42h23.5A6.25 6.25 0 0 0 42 35.75V16.5H6Zm12 7a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm6 2.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm11-2.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM15.5 34a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm11-2.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`),
		g.Group(children),
	)
}