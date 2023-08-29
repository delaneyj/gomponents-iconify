package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWatch0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M34.956 31L31 44H17l-3.957-13m0-14L17 4h14l3.956 13"/><path fill="#fff" stroke="#fff" d="M37 24c0 2.577-.75 4.98-2.044 7A12.99 12.99 0 0 1 24 37a12.99 12.99 0 0 1-10.956-6A12.94 12.94 0 0 1 11 24c0-2.577.75-4.98 2.044-7A12.99 12.99 0 0 1 24 11a12.99 12.99 0 0 1 10.956 6A12.94 12.94 0 0 1 37 24Z"/><path stroke="#000" d="M24 17v7l4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWatch0)"/>`),
		g.Group(children),
	)
}