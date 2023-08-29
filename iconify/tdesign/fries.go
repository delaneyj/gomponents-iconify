package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.465.994L21 4.53l-.707.707l2.122 2.121l-2.122 2.121l2.122 2.122l-3.2 3.2l1.397 1.397L8.845 23.26L.741 15.155l7.06-11.768l1.514 1.514L13.222.994l2.122 2.121l2.12-2.12ZM12.85 8.437l.707.707l4.614-4.614l-.707-.707l-4.614 4.614Zm6.028-1.786l-3.907 3.907l.707.707l3.907-3.907l-.707-.707Zm-1.786 6.028l.707.707l1.786-1.785l-.707-.707l-1.786 1.785Zm.294 3.122L8.198 6.613l-1.94 3.232l7.897 7.896l3.232-1.94Zm-5 3l-7.189-7.188l-1.939 3.232l5.896 5.896l3.232-1.94ZM10.729 6.315l.708.707l2.492-2.492l-.707-.707l-2.493 2.492Z"/>`),
		g.Group(children),
	)
}