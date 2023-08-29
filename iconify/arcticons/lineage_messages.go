package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageMessages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.201 12.593l1.15-5.865a2.115 2.115 0 0 0-2.73-2.418l-4.99 1.625a21.578 21.578 0 1 0 6.57 6.658Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 12.316c9.273 0 11.944 8.926 11.944 14.203S32.057 35.64 24 35.64s-11.944-3.844-11.944-9.121S14.726 12.316 24 12.316Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.666c3.324 0 6.284-1.629 7.5-5.798a16.544 16.544 0 0 1-7.5 1.368a16.544 16.544 0 0 1-7.5-1.368c1.216 4.17 4.176 5.798 7.5 5.798Z"/>`),
		g.Group(children),
	)
}