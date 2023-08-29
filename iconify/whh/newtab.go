package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newtab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.214 1024h-768q-53 0-90.5-37.5T.214 896V256l48-156q14-43 58.5-71.5t95.5-28.5h364q51 0 95.5 28.5t58.5 71.5l48 156h128q53 0 90.5 37.5t37.5 90.5v512q0 53-37.5 90.5t-90.5 37.5zm0-608q0-13-9.5-22.5t-22.5-9.5h-704q-13 0-22.5 9.5t-9.5 22.5v448q0 13 9.5 22.5t22.5 9.5h704q13 0 22.5-9.5t9.5-22.5V416zm-256 288h-64v64q0 26-19 45t-45 19t-45-19t-19-45v-64h-64q-27 0-45.5-18.5t-18.5-45.5t18.5-45.5t45.5-18.5h64v-64q0-26 19-45t45-19t45 18.5t19 45.5v64h64q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}