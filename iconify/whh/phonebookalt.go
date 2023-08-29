package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonebookalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.356 960h-64V704h64q27 0 45.5 18.5t18.5 45.5v128q0 27-18.5 45.5t-45.5 18.5zm0-320h-64V384h64q27 0 45.5 18.5t18.5 45.5v128q0 27-18.5 45.5t-45.5 18.5zm0-320h-64V64h64q27 0 45.5 18.5t18.5 45.5v128q0 27-18.5 45.5t-45.5 18.5zm-128 640q0 26-18.5 45t-45.5 19h-704q-26 0-45-19t-19-45V64q0-27 19-45.5t45-18.5h704q27 0 45.5 18.5t18.5 45.5v896zm-352-346v-18q96-39 96-212q0-63-44.5-95.5t-115-32.5t-115.5 32.5t-45 95.5q0 168 96 211v19q-69 12-114 45t-46 75q0 12 21.5 19.5t63 10.5t68.5 3.5t71 .5t71-.5t68.5-3.5t63-10.5t21.5-19.5q-1-42-46-75t-114-45z"/>`),
		g.Group(children),
	)
}