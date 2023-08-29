package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.66 7H15v1a1 1 0 0 0 2 0V7h1a1 1 0 0 1 1 1v3h-1.34a1 1 0 0 0 0 2H19v1.34a1 1 0 1 0 2 0V8a3 3 0 0 0-3-3h-1V4a1 1 0 0 0-2 0v1h-3.34a1 1 0 0 0 0 2Zm10.05 13.29l-1.6-1.6l-16.4-16.4a1 1 0 0 0-1.42 1.42l1.91 1.9A3 3 0 0 0 3 8v10a3 3 0 0 0 3 3h12a3 3 0 0 0 1.29-.3l1 1a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41ZM5 8a1 1 0 0 1 .66-.93L9.59 11H5Zm1 11a1 1 0 0 1-1-1v-5h6.59l6 6Z"/>`),
		g.Group(children),
	)
}