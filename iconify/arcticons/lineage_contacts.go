package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageContacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.584 16.578c10.158 0 13.084 9.777 13.084 15.558s-4.258 9.991-13.084 9.991S5.5 37.917 5.5 32.136s2.926-15.558 13.084-15.558Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.584 37.774c3.642 0 6.883-1.784 8.216-6.352a18.122 18.122 0 0 1-8.216 1.499a18.122 18.122 0 0 1-8.216-1.499c1.333 4.568 4.575 6.352 8.216 6.352ZM16.92 16.672c1.277-5.221 4.823-10.8 12.496-10.8c10.158 0 13.084 9.778 13.084 15.559c0 5.267-3.536 9.231-10.84 9.894"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.97 26.953c2.988-.46 5.523-2.33 6.662-6.236a18.122 18.122 0 0 1-8.216 1.499q-.237 0-.466-.002"/>`),
		g.Group(children),
	)
}