package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdmissionAssistant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.728 43.5h26.239L24 14.112m-4.003 6.933L7.033 43.5m20.894-8.719l-7.93-13.736m0 13.736l3.965-6.868m3.965 6.868h-7.93M13.691 7.854L24.132 4.5l9.37 3.354l-4.114 1.66v4.598a9.522 9.522 0 0 0-5.72-1.625c-3.719 0-4.998 1.383-4.998 1.383V9.86Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.303 8.891a19.262 19.262 0 0 0-4.633.968m-2.901-1.168v4.625"/>`),
		g.Group(children),
	)
}