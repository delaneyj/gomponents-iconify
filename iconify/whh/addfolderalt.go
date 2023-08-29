package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addfolderalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.94 896h-896q-27 0-45.5-19T.94 832V256h1024v576q0 26-19 45t-45 19zm-64-352q0-13-9.5-22.5t-22.5-9.5h-96v-96q0-13-9.5-22.5t-22.5-9.5h-64q-13 0-22.5 9.5t-9.5 22.5v96h-96q-13 0-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5h96v96q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5v-96h96q13 0 22.5-9.5t9.5-22.5v-64zM.94 64q0-26 18.5-45t45.5-19h384q27 0 45.5 18.5t18.5 45.5t18.5 45.5t45.5 18.5h384q26 0 45 19t19 45H.94V64z"/>`),
		g.Group(children),
	)
}