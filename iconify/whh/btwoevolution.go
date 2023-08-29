package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Btwoevolution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992.232 320h-32v128l-128-128v448q0 27-18.5 45.5t-45.5 18.5h-64v192l-192-192h-256q-27 0-45.5-18.5t-18.5-45.5V576h-128q-27 0-45.5-18.5T.232 512V128q0-27 18.5-45.5t45.5-18.5h384q27 0 45.5 18.5t18.5 45.5v64h128V32q0-13 9.5-22.5t22.5-9.5h320q13 0 22.5 9.5t9.5 22.5v256q0 13-9.5 22.5t-22.5 9.5zm-544-160q0-13-9.5-22.5t-22.5-9.5h-320q-13 0-22.5 9.5t-9.5 22.5v320q0 13 9.5 22.5t22.5 9.5h96V256q0-26 18.5-45t45.5-19h192v-32zm512-64q0-13-9.5-22.5t-22.5-9.5h-192q-13 0-22.5 9.5t-9.5 22.5v96h64q26 0 45 18.5t19 45.5h96q13 0 22.5-9.5t9.5-22.5V96z"/>`),
		g.Group(children),
	)
}