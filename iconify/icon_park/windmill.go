package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windmill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"><path d="M24 23.9917L23.9707 13.9958L23.9415 4L12 14V24L24 23.9917Z"/><path d="M24.0083 24L34.0042 23.9707L44 23.9415L34 12L24 12L24.0083 24Z"/><path d="M24 24.0083L24.0293 34.0042L24.0585 44L36 34V24L24 24.0083Z"/><path d="M23.9917 24L13.9958 24.0293L4 24.0585L14 36L24 36L23.9917 24Z"/></g>`),
		g.Group(children),
	)
}