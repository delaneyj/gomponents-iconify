package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skrill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-236q101-44 168.5-129.5t67.5-179.5q0-66-24-124t-70-104t-121-73t-169-28v-1q-20 0-32.5-.5t-32.5-4t-32-9.5t-21.5-19t-9.5-31q0-23 21.5-38.5t48.5-20.5t58-5q193 0 320 64V64q-37-37-131-64h195q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-768-714q0 66 24 124.5t70 104.5t121 73t169 28q20 0 32.5.5t32.5 4t32 9.5t21.5 19t9.5 31q0 23-21.5 38.5t-48.5 20.5t-58 5q-193 0-320-64v256q37 37 131 64h-195q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h236q-101 44-168.5 130t-67.5 180z"/>`),
		g.Group(children),
	)
}