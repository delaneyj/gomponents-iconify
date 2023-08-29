package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squaregithub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-258q2-13 2-25q0-85-19.5-170.5t-54.5-127.5q130-11 198-76.5t68-169.5q0-91-64-140v-91q0-10-11-21t-21-11q-12 0-119 70q-47-6-105-6t-105 6q-107-70-119-70q-10 0-21 11t-11 21q0 9-2 44t-3 51q-59 49-59 136q0 104 68 169.5t198 76.5q-5 7-14 21q-43 46-124 46q-16 0-32-13t-31.5-32t-32.5-38t-42-32t-54-13q0 8 10 20.5t26.5 33t27.5 42.5q26 51 54 73.5t74 22.5q48 0 87-14q-23 88-23 181q0 12 2 25h-258q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5z"/>`),
		g.Group(children),
	)
}