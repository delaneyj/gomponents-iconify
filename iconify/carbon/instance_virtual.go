package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceVirtual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="7" cy="23" r="1" fill="currentColor"/><path fill="currentColor" d="M2 6h4v2H2zm6 0h4v2H8zm6 0h4v2h-4zm6 0h4v2h-4zm6 0h4v2h-4zm2 22H4a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zM4 20v6h24v-6zm-2-8h28v2H2z"/>`),
		g.Group(children),
	)
}