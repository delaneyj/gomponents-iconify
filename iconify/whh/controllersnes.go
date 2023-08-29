package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Controllersnes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 568q-96 0-169-64H425q-73 64-169 64q-106 0-181-75T0 312t75-181t181-75h512q106 0 181 75t75 181t-75 181t-181 75zM384 248h-64v-64q0-27-19-45.5T255.5 120t-45 18.5T192 184v64h-64q-27 0-45.5 19T64 312t18.5 45t45.5 19h64v64q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5v-64h64q26 0 45-19t19-45.5t-19-45t-45-18.5zm319.5 64q-26.5 0-45 19T640 376.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19zm128-128q-26.5 0-45 19T768 248.5t18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19z"/>`),
		g.Group(children),
	)
}