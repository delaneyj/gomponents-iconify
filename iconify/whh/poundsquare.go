package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poundsquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-448-448h64q26 0 45-18.5t19-45t-19-45.5t-45-19h-112l-2-6q-14-28-14-58q0-53 37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5q0 26 18.5 45t45 19t45.5-19t19-45q0-106-75-181t-181-75t-181 75t-75 181q0 29 9 64h-9q-27 0-45.5 19t-18.5 45t18.5 45t45.5 19h64q0 35-10 66t-22 49t-22 39t-10 38q0 26 18.5 45t45.5 19h384q26 0 45-18.5t19-45.5t-19-45.5t-45-18.5h-292q16-15 24.5-47t10.5-57z"/>`),
		g.Group(children),
	)
}