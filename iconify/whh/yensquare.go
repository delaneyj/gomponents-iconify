package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yensquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-81.5-815q-18.5-18-44-18t-43.5 18l-215 163l-216-163q-18-18-43.5-18t-43.5 18t-18 43.5t18 43.5l239 181v35h-96q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h96v64h-96q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h96v128q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5V704h96q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-96v-64h96q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-96v-35l239-181q18-18 18-43.5t-18.5-43.5z"/>`),
		g.Group(children),
	)
}