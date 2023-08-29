package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaskOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.42 19.41A2 2 0 0 1 18 20H6a2 2 0 0 1-2-2V6c0-.554.225-1.055.588-1.417M8 4h10a2 2 0 0 1 2 2v10"/><path d="M9.885 9.872a3 3 0 1 0 4.245 4.24m.582-3.396a3.012 3.012 0 0 0-1.438-1.433M3 3l18 18"/></g>`),
		g.Group(children),
	)
}