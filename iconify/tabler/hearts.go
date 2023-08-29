package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.017 18L12 20l-7.5-7.428A5 5 0 1 1 12 6.006a5 5 0 0 1 8.153 5.784"/><path d="m15.99 20l4.197-4.223a2.81 2.81 0 0 0 0-3.948a2.747 2.747 0 0 0-3.91-.007l-.28.282l-.279-.283a2.747 2.747 0 0 0-3.91-.007a2.81 2.81 0 0 0-.007 3.948L15.983 20z"/></g>`),
		g.Group(children),
	)
}