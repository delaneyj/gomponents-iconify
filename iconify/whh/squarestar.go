package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarestar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-80-620l-186-40l-96-163q-9-9-21.5-9t-21.5 9l-97 163l-186 40q-11 6-14.5 18t1.5 23l129 144l-15 182q2 13 9.5 22t19.5 7l174-76l175 76q11 2 19-7t10-22l-15-182l128-144q6-11 2-23t-15-18z"/>`),
		g.Group(children),
	)
}