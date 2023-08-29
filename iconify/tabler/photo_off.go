package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h.01M7 3h11a3 3 0 0 1 3 3v11m-.856 3.099A2.991 2.991 0 0 1 18 21H6a3 3 0 0 1-3-3V6c0-.845.349-1.608.91-2.153"/><path d="m3 16l5-5c.928-.893 2.072-.893 3 0l5 5m.33-3.662c.574-.054 1.155.166 1.67.662l3 3M3 3l18 18"/></g>`),
		g.Group(children),
	)
}