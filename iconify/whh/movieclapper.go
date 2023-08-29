package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Movieclapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M969.057 248h-1l-71-119l-127 17l80 120l-212 33l-81-120l-116 18l81 120l-220 35l-77-127l-122 24l81 121l-95 15h104l-64 127h128l64-127h192l-64 127h128l64-127h192l-64 127h128l64-127q26 0 45 18.5t19 44.5v512q0 27-19 45.5t-45 18.5h-895q-27 0-45.5-18.5T2.057 960V448q0-26 18.5-44.5t45.5-18.5q-5-2-12-5.5t-20.5-17.5t-16.5-31l-16-124q-3-25 13.5-45.5t42.5-23.5l882-137q26-3 46.5 13t23.5 41l16 124q3 25-13.5 45.5t-42.5 23.5zm-808 648h704q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-704q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5zm0-192h704q13 0 22.5-9t9.5-22.5t-9.5-23t-22.5-9.5h-704q-13 0-22.5 9.5t-9.5 23t9.5 22.5t22.5 9z"/>`),
		g.Group(children),
	)
}