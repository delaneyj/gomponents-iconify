package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 21h14v2H9zm2-7a2 2 0 1 0 2 2a1.98 1.98 0 0 0-2-2zm10 0a2 2 0 1 0 2 2a1.98 1.98 0 0 0-2-2z"/><path fill="currentColor" d="M28 8h-9.586l2.072-2.072A2.04 2.04 0 0 0 21 6a2 2 0 1 0-2-2a2.041 2.041 0 0 0 .072.514L15.586 8H4a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V10a2.002 2.002 0 0 0-2-2ZM4 26V10h24v16Z"/>`),
		g.Group(children),
	)
}