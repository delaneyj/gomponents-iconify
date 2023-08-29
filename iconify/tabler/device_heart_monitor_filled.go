package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceHeartMonitorFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18 3a3 3 0 0 1 2.995 2.824L21 6v12a3 3 0 0 1-2.824 2.995L18 21H6a3 3 0 0 1-2.995-2.824L3 18V6a3 3 0 0 1 2.824-2.995L6 3h12zm-4 13a1 1 0 0 0-.993.883L13 17l.007.127a1 1 0 0 0 1.986 0L15 17.01l-.007-.127A1 1 0 0 0 14 16zm3 0a1 1 0 0 0-.993.883L16 17l.007.127a1 1 0 0 0 1.986 0L18 17.01l-.007-.127A1 1 0 0 0 17 16zm-6-6.764l-.106.211a1 1 0 0 1-.77.545L10 10l-5-.001V13h14V9.999L14.618 10l-.724 1.447a1 1 0 0 1-1.725.11l-.063-.11L11 9.236zM18 5H6a1 1 0 0 0-.993.883L5 6v1.999L9.381 8l.725-1.447a1 1 0 0 1 1.725-.11l.063.11L13 8.763l.106-.21a1 1 0 0 1 .77-.545L14 8l5-.001V6a1 1 0 0 0-.883-.993L18 5z"/></g>`),
		g.Group(children),
	)
}