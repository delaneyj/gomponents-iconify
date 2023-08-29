package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 5.27a8.5 8.5 0 0 1 5.09 1.686L18.807 8.74l8.428 2.255l-2.26-8.427l-1.89 1.89A12.01 12.01 0 0 0 15.5 1.77C8.827 1.773 3.418 7.18 3.417 13.855c0 4.063 2.012 7.647 5.084 9.838v-4.887a8.55 8.55 0 0 1-1.584-4.952a8.594 8.594 0 0 1 8.584-8.584zm-6 23.96h12V12.355h-12V29.23z"/>`),
		g.Group(children),
	)
}