package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whatsappbusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.95 24a6 6 0 1 1 0 12h-9.9V12h9.9a6 6 0 1 1 0 12Zm-.081 0h-9.488"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.517 21.517 0 0 0 5.148 34.36L2.5 45.5l11.14-2.648A21.504 21.504 0 1 0 24 2.5Z"/>`),
		g.Group(children),
	)
}