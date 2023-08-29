package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M35.5 4.5c0-2.38-.5-3-3-3h-25c-2.5 0-3 .561-3 3v34c0 2.41.561 3 3 3h25c2.471 0 3-.561 3-3v-34zm-26 2h21v27h-21v-27zm14 32h-7v-2h7v2z"/>`),
		g.Group(children),
	)
}