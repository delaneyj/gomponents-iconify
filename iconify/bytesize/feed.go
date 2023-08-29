package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="26" r="2" fill="currentColor"/><path d="M4 15c7 0 13 6 13 13M4 6c13 0 22 9 22 22"/></g>`),
		g.Group(children),
	)
}