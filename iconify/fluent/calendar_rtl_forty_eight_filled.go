package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M42 12.25A6.25 6.25 0 0 0 35.75 6h-23.5A6.25 6.25 0 0 0 6 12.25V14h36v-1.75Zm0 4.25v19.25A6.25 6.25 0 0 1 35.75 42h-23.5A6.25 6.25 0 0 1 6 35.75V16.5h36Zm-12 7a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0ZM24 26a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-11-2.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0ZM32.5 34a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-11-2.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Z"/>`),
		g.Group(children),
	)
}